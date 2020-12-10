// brian taylor vann
// copycopy
const copycopy = (atomToCopy) => {
    if (atomToCopy instanceof Object === false) {
        return atomToCopy;
    }
    const entries = Array.isArray(atomToCopy)
        ? [...atomToCopy]
        : Object.assign({}, atomToCopy);
    for (const index in entries) {
        const entry = entries[index];
        if (entries instanceof Object) {
            entries[index] = copycopy(entry);
        }
    }
    return entries;
};

const buildResultsState = ({ testCollection, startTime, stub, }) => {
    const nextState = {
        status: "submitted",
        results: [],
        startTime,
    };
    for (const collection of testCollection) {
        const { tests, title } = collection;
        const collectionResults = {
            title,
            status: "unsubmitted",
        };
        const results = [];
        for (const test of tests) {
            const { name } = test;
            results.push({
                status: "unsubmitted",
                name,
            });
        }
        if (nextState.results) {
            nextState.results.push(Object.assign(Object.assign({}, collectionResults), { results }));
        }
    }
    return nextState;
};

const startTestCollectionState = (runResults, params) => {
    var _a;
    if (runResults.results === undefined) {
        return runResults;
    }
    const { startTime, collectionID } = params;
    const collectionResult = (_a = runResults === null || runResults === void 0 ? void 0 : runResults.results) === null || _a === void 0 ? void 0 : _a[collectionID];
    if (collectionResult) {
        collectionResult.status = "submitted";
        collectionResult.startTime = startTime;
    }
    return runResults;
};

const startTestState = (runResults, params) => {
    var _a, _b, _c;
    if (runResults.results === undefined) {
        return runResults;
    }
    const { startTime, collectionID, testID } = params;
    const testResult = (_c = (_b = (_a = runResults === null || runResults === void 0 ? void 0 : runResults.results) === null || _a === void 0 ? void 0 : _a[collectionID]) === null || _b === void 0 ? void 0 : _b.results) === null || _c === void 0 ? void 0 : _c[testID];
    if (testResult) {
        testResult.status = "submitted";
        testResult.startTime = startTime;
    }
    return runResults;
};

const cancelRunState = (runResults, params) => {
    const { endTime } = params;
    runResults.endTime = endTime;
    runResults.status = "cancelled";
    const collectionResults = runResults.results;
    if (collectionResults) {
        for (const collection of collectionResults) {
            if (collection.status === "submitted") {
                collection.status = "cancelled";
            }
            const testResults = collection.results;
            if (testResults) {
                for (const result of testResults) {
                    if (result.status === "submitted") {
                        result.status = "cancelled";
                    }
                }
            }
        }
    }
    return runResults;
};

const allTestsHavePassed = (testResults) => {
    for (const result of testResults) {
        if (result.status !== "passed") {
            return false;
        }
    }
    return true;
};
const endTestCollectionState = (runResults, params) => {
    if (runResults.results === undefined) {
        return runResults;
    }
    const { endTime, collectionID } = params;
    const collection = runResults.results[collectionID];
    if (collection === undefined) {
        return runResults;
    }
    collection.endTime = endTime;
    collection.status = "failed";
    const collectionResults = collection.results;
    if (collectionResults && allTestsHavePassed(collectionResults)) {
        collection.status = "passed";
    }
    return runResults;
};

const endTestState = (runResults, params) => {
    var _a, _b, _c;
    if (runResults.results === undefined) {
        return runResults;
    }
    const { assertions, endTime, collectionID, testID } = params;
    const testResult = (_c = (_b = (_a = runResults === null || runResults === void 0 ? void 0 : runResults.results) === null || _a === void 0 ? void 0 : _a[collectionID]) === null || _b === void 0 ? void 0 : _b.results) === null || _c === void 0 ? void 0 : _c[testID];
    if (testResult === undefined) {
        return runResults;
    }
    testResult.status = "failed";
    if (assertions === undefined) {
        testResult.status = "passed";
    }
    if (assertions && assertions.length === 0) {
        testResult.status = "passed";
    }
    testResult.assertions = assertions;
    testResult.endTime = endTime;
    return runResults;
};

// for test collection
const allTestCollectionsHavePassed = (collectionResults) => {
    for (const collection of collectionResults) {
        if (collection.status === "failed") {
            return false;
        }
    }
    return true;
};
const endTestRunState = (runResults, params) => {
    const { endTime } = params;
    runResults.endTime = endTime;
    runResults.status = "failed";
    const results = runResults.results;
    if (results && allTestCollectionsHavePassed(results)) {
        runResults.status = "passed";
    }
    return runResults;
};

// brian taylor vann
const defaultResultsState = {
    status: "unsubmitted",
};
let resultsState = Object.assign({}, defaultResultsState);
const buildResults = (params) => {
    resultsState = buildResultsState(params);
};
const startTestCollection = (params) => {
    const copyOfResults = copycopy(resultsState);
    resultsState = startTestCollectionState(copyOfResults, params);
};
const startTest = (params) => {
    const copyOfResults = copycopy(resultsState);
    resultsState = startTestState(copyOfResults, params);
};
const cancelRun = (params) => {
    const copyOfResults = copycopy(resultsState);
    resultsState = cancelRunState(copyOfResults, params);
};
const endTest = (params) => {
    const copyOfResults = copycopy(resultsState);
    resultsState = endTestState(copyOfResults, params);
};
const endTestCollection = (params) => {
    const copyOfResults = copycopy(resultsState);
    resultsState = endTestCollectionState(copyOfResults, params);
};
const endTestRun = (params) => {
    const copyOfResults = copycopy(resultsState);
    resultsState = endTestRunState(copyOfResults, params);
};
const getResults = () => {
    return copycopy(resultsState);
};

// brian taylor vann
class SubPub {
    constructor() {
        this.stub = 0;
        this.recycledStubs = [];
        this.subscriptions = {};
    }
    getStub() {
        const stub = this.recycledStubs.pop();
        if (stub) {
            return stub;
        }
        this.stub += 1;
        return this.stub;
    }
    subscribe(callback) {
        const stub = this.getStub();
        this.subscriptions[stub] = callback;
        return stub;
    }
    unsubscribe(stub) {
        this.subscriptions[stub] = undefined;
        this.recycledStubs.push(stub);
    }
    broadcast(params) {
        for (const stubKey in this.subscriptions) {
            const subscription = this.subscriptions[stubKey];
            if (subscription !== undefined) {
                subscription(params);
            }
        }
    }
}

const subpub = new SubPub();
// send current state to subscribers
const broadcast = (testRunState) => {
    subpub.broadcast(testRunState);
};

// little test runner
const START_TEST_RUN = "START_TEST_RUN";
const START_TEST_COLLECTION = "START_TEST_COLLECTION";
const START_TEST = "START_TEST";
const CANCEL_RUN = "CANCEL_RUN";
const END_TEST = "END_TEST";
const END_TEST_COLLECTION = "END_TEST_COLLECTION";
const END_TEST_RUN = "END_TEST_RUN";
const consolidate = (action) => {
    switch (action.action) {
        case START_TEST_RUN:
            buildResults(action.params);
            break;
        case START_TEST_COLLECTION:
            startTestCollection(action.params);
            break;
        case START_TEST:
            startTest(action.params);
            break;
        case CANCEL_RUN:
            cancelRun(action.params);
            break;
        case END_TEST:
            endTest(action.params);
            break;
        case END_TEST_COLLECTION:
            endTestCollection(action.params);
            break;
        case END_TEST_RUN:
            endTestRun(action.params);
            break;
    }
    broadcast(getResults());
};
const dispatch = (action) => {
    consolidate(action);
};

// brian taylor vann
// run tests
const startTestRun = (params) => {
    dispatch({
        action: "START_TEST_RUN",
        params,
    });
};
const startTestCollection$1 = (params) => {
    dispatch({
        action: "START_TEST_COLLECTION",
        params,
    });
};
const startTest$1 = (params) => {
    dispatch({
        action: "START_TEST",
        params,
    });
};
const sendTestResult = (params) => {
    dispatch({
        action: "END_TEST",
        params,
    });
};
const endTestCollection$1 = (params) => {
    dispatch({
        action: "END_TEST_COLLECTION",
        params,
    });
};
const endTestRun$1 = (params) => {
    dispatch({
        action: "END_TEST_RUN",
        params,
    });
};

// brian taylor vann
let stub = 0;
const getStub = () => {
    return stub;
};
const updateStub = () => {
    stub += 1;
    return stub;
};

// little test runner
// brian taylor vann
var __awaiter = (undefined && undefined.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
const sleep = (time) => __awaiter(void 0, void 0, void 0, function* () {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve();
        }, time);
    });
});
const defaultTimeoutInterval = 10000;
const getTimeoutAssertions = (timeoutInterval) => [
    `timed out at: ${timeoutInterval}`,
];
const createTestTimeout = (timeoutInterval) => __awaiter(void 0, void 0, void 0, function* () {
    const interval = timeoutInterval !== null && timeoutInterval !== void 0 ? timeoutInterval : defaultTimeoutInterval;
    yield sleep(interval);
    return getTimeoutAssertions(interval);
});
const buildTest = (params) => {
    const { issuedAt, testID, collectionID, timeoutInterval } = params;
    return () => __awaiter(void 0, void 0, void 0, function* () {
        if (issuedAt < getStub()) {
            return;
        }
        const startTime = performance.now();
        startTest$1({
            collectionID,
            testID,
            startTime,
        });
        const assertions = yield Promise.race([
            params.testFunc(),
            createTestTimeout(timeoutInterval),
        ]);
        if (issuedAt < getStub()) {
            return;
        }
        const endTime = performance.now();
        sendTestResult({
            endTime,
            assertions,
            collectionID,
            testID,
        });
    });
};
const runTestsAllAtOnce = ({ startTime, collectionID, tests, timeoutInterval, }) => __awaiter(void 0, void 0, void 0, function* () {
    const builtAsyncTests = [];
    let testID = 0;
    for (const testFunc of tests) {
        builtAsyncTests.push(buildTest({
            collectionID,
            issuedAt: startTime,
            testFunc,
            testID,
            timeoutInterval,
        })() // execute test before push
        );
        testID += 1;
    }
    if (startTime < getStub()) {
        return;
    }
    yield Promise.all(builtAsyncTests);
});
const runTestsInOrder = ({ startTime, collectionID, tests, timeoutInterval, }) => __awaiter(void 0, void 0, void 0, function* () {
    let testID = 0;
    for (const testFunc of tests) {
        if (startTime < getStub()) {
            return;
        }
        const builtTest = buildTest({
            collectionID,
            issuedAt: startTime,
            testFunc,
            testID,
            timeoutInterval,
        });
        yield builtTest();
        testID += 1;
    }
});

// little test runner
// brian taylor vann
var __awaiter$1 = (undefined && undefined.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
// create a test collection
const startLtrTestCollectionRun = ({ testCollection, startTime, stub, }) => __awaiter$1(void 0, void 0, void 0, function* () {
    startTestRun({ testCollection, startTime, stub });
    let collectionID = 0;
    for (const collection of testCollection) {
        if (stub < getStub()) {
            return;
        }
        const { tests, runTestsAsynchronously, timeoutInterval } = collection;
        const runParams = {
            collectionID,
            tests,
            startTime,
            timeoutInterval,
        };
        startTestCollection$1({
            collectionID,
            startTime,
        });
        if (runTestsAsynchronously) {
            yield runTestsAllAtOnce(runParams);
        }
        else {
            yield runTestsInOrder(runParams);
        }
        if (stub < getStub()) {
            return;
        }
        const endTime = performance.now();
        endTestCollection$1({
            collectionID,
            endTime,
        });
        collectionID += 1;
    }
    if (stub < getStub()) {
        return;
    }
    const endTime = performance.now();
    endTestRun$1({ endTime });
});
// iterate through tests synchronously
const runTests = (params) => __awaiter$1(void 0, void 0, void 0, function* () {
    const startTime = performance.now();
    const stub = updateStub();
    yield startLtrTestCollectionRun(Object.assign(Object.assign({}, params), { startTime, stub }));
    if (startTime < getStub()) {
        return;
    }
    return getResults();
});

// brian taylor vann
const createAlphabetKeys = (route) => {
    const alphabetSet = {};
    let lowercaseIndex = "a".charCodeAt(0);
    const lowercaseLimit = "z".charCodeAt(0);
    let uppercaseIndex = "A".charCodeAt(0);
    const uppercaseLimit = "Z".charCodeAt(0);
    while (lowercaseIndex <= lowercaseLimit) {
        alphabetSet[String.fromCharCode(lowercaseIndex)] = route;
        lowercaseIndex += 1;
    }
    while (uppercaseIndex <= uppercaseLimit) {
        alphabetSet[String.fromCharCode(uppercaseIndex)] = route;
        uppercaseIndex += 1;
    }
    return alphabetSet;
};
const routers = {
    CONTENT_NODE: {
        "<": "OPEN_NODE",
        DEFAULT: "CONTENT_NODE",
    },
    OPEN_NODE: Object.assign(Object.assign({}, createAlphabetKeys("OPEN_NODE_VALID")), { "<": "OPEN_NODE", "/": "CLOSE_NODE", DEFAULT: "CONTENT_NODE" }),
    OPEN_NODE_VALID: {
        "<": "OPEN_NODE",
        "/": "INDEPENDENT_NODE_VALID",
        ">": "OPEN_NODE_CONFIRMED",
        DEFAULT: "OPEN_NODE_VALID",
    },
    CLOSE_NODE: Object.assign(Object.assign({}, createAlphabetKeys("CLOSE_NODE_VALID")), { "<": "OPEN_NODE", DEFAULT: "CONTENT_NODE" }),
    CLOSE_NODE_VALID: {
        "<": "OPEN_NODE",
        ">": "CLOSE_NODE_CONFIRMED",
        DEFAULT: "CLOSE_NODE_VALID",
    },
    INDEPENDENT_NODE_VALID: {
        "<": "OPEN_NODE",
        ">": "INDEPENDENT_NODE_CONFIRMED",
        DEFAULT: "INDEPENDENT_NODE_VALID",
    },
};

const DEFAULT_POSITION = {
    arrayIndex: 0,
    stringIndex: 0,
};
const create = (position = DEFAULT_POSITION) => ({
    origin: Object.assign({}, position),
    target: Object.assign({}, position),
});
const createFollowingVector = (vector, template) => {
    const followingVector = copy(vector);
    increment(followingVector.target, template);
    followingVector.origin = Object.assign({}, followingVector.target);
    return followingVector;
};
const copy = (vector) => {
    return {
        origin: Object.assign({}, vector.origin),
        target: Object.assign({}, vector.target),
    };
};
const increment = (position, template) => {
    // template boundaries
    const templateLength = template.templateArray.length;
    const chunkLength = template.templateArray[position.arrayIndex].length;
    if (chunkLength === undefined) {
        return;
    }
    // determine if finished
    if (position.arrayIndex >= templateLength - 1 && chunkLength === 0) {
        return;
    }
    if (position.arrayIndex >= templateLength - 1 &&
        position.stringIndex >= chunkLength - 1) {
        return;
    }
    // cannot % modulo by 0
    if (chunkLength > 0) {
        position.stringIndex += 1;
        position.stringIndex %= chunkLength;
    }
    if (position.stringIndex === 0) {
        position.arrayIndex += 1;
    }
    return position;
};
const getCharFromTarget = (vector, template) => {
    const templateArray = template.templateArray;
    const arrayIndex = vector.target.arrayIndex;
    const stringIndex = vector.target.stringIndex;
    if (arrayIndex > templateArray.length - 1) {
        return;
    }
    if (stringIndex > templateArray[arrayIndex].length - 1) {
        return;
    }
    return templateArray[arrayIndex][stringIndex];
};

// brian taylor vann
const DEFAULT = "DEFAULT";
const CONTENT_NODE = "CONTENT_NODE";
const OPEN_NODE = "OPEN_NODE";
const validSieve = {
    ["OPEN_NODE_VALID"]: "OPEN_NODE_VALID",
    ["CLOSE_NODE_VALID"]: "CLOSE_NODE_VALID",
    ["INDEPENDENT_NODE_VALID"]: "INDEPENDENT_NODE_VALID",
};
const confirmedSieve = {
    ["OPEN_NODE_CONFIRMED"]: "OPEN_NODE_CONFIRMED",
    ["CLOSE_NODE_CONFIRMED"]: "CLOSE_NODE_CONFIRMED",
    ["INDEPENDENT_NODE_CONFIRMED"]: "INDEPENDENT_NODE_CONFIRMED",
};
const setStartStateProperties = (template, previousCrawl) => {
    const cState = {
        nodeType: CONTENT_NODE,
        vector: create(),
    };
    if (previousCrawl !== undefined) {
        cState.vector = createFollowingVector(previousCrawl.vector, template);
    }
    setNodeType(template, cState);
    return cState;
};
const setNodeType = (template, cState) => {
    var _a, _b;
    const nodeStates = routers[cState.nodeType];
    const char = getCharFromTarget(cState.vector, template);
    if (nodeStates !== undefined && char !== undefined) {
        const defaultNodeType = (_a = nodeStates[DEFAULT]) !== null && _a !== void 0 ? _a : CONTENT_NODE;
        cState.nodeType = (_b = nodeStates[char]) !== null && _b !== void 0 ? _b : defaultNodeType;
    }
    return cState;
};
const crawl = (template, previousCrawl) => {
    let openPosition;
    const cState = setStartStateProperties(template, previousCrawl);
    while (increment(cState.vector.target, template)) {
        if (cState.vector.target.stringIndex === 0) {
            console.log(template.templateArray[cState.vector.target.arrayIndex]);
        }
        if (validSieve[cState.nodeType] === undefined &&
            cState.vector.target.stringIndex === 0) {
            cState.nodeType = CONTENT_NODE;
        }
        setNodeType(template, cState);
        if (confirmedSieve[cState.nodeType]) {
            if (openPosition !== undefined) {
                cState.vector.origin = Object.assign({}, openPosition);
            }
            break;
        }
        if (cState.nodeType === OPEN_NODE) {
            openPosition = Object.assign({}, cState.vector.target);
        }
    }
    return cState;
};

// brian taylor vann
const testTextInterpolator = (templateArray, ...injections) => {
    return { templateArray, injections };
};
const title = "crawl";
const runTestsAsynchronously = true;
const findNothingWhenThereIsPlainText = () => {
    const testBlank = testTextInterpolator `no nodes to be found!`;
    const assertions = [];
    const result = crawl(testBlank);
    if (result === undefined) {
        assertions.push("undefined result");
    }
    if (result && result.nodeType !== "CONTENT_NODE") {
        assertions.push(`should return CONTENT_NODE instead of ${result.nodeType}`);
    }
    if (result && result.vector.origin.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.vector.origin.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.vector.target.arrayIndex !== 0) {
        assertions.push(`should return end arrayIndex as 0`);
    }
    if (result && result.vector.target.stringIndex !== 20) {
        assertions.push(`should return end stringIndex as 20`);
    }
    return assertions;
};
const findParagraphInPlainText = () => {
    const testOpenNode = testTextInterpolator `<p>`;
    const assertions = [];
    const result = crawl(testOpenNode);
    if (result === undefined) {
        assertions.push("undefined result");
    }
    if (result && result.nodeType !== "OPEN_NODE_CONFIRMED") {
        assertions.push(`should return OPEN_NODE_CONFIRMED instead of ${result.nodeType}`);
    }
    if (result && result.vector.origin.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.vector.origin.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.vector.target.arrayIndex !== 0) {
        assertions.push(`should return end arrayIndex as 0`);
    }
    if (result && result.vector.target.stringIndex !== 2) {
        assertions.push(`should return end stringIndex as 2`);
    }
    return assertions;
};
const findCloseParagraphInPlainText = () => {
    const testTextCloseNode = testTextInterpolator `</p>`;
    const assertions = [];
    const result = crawl(testTextCloseNode);
    if (result === undefined) {
        assertions.push("undefined result");
    }
    if (result && result.nodeType !== "CLOSE_NODE_CONFIRMED") {
        assertions.push(`should return CLOSE_NODE_CONFIRMED instead of ${result.nodeType}`);
    }
    if (result && result.vector.origin.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.vector.origin.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 2`);
    }
    if (result && result.vector.target.arrayIndex !== 0) {
        assertions.push(`should return end arrayIndex as 0`);
    }
    if (result && result.vector.target.stringIndex !== 3) {
        assertions.push(`should return end stringIndex as 3`);
    }
    return assertions;
};
const findIndependentParagraphInPlainText = () => {
    const testTextIndependentNode = testTextInterpolator `<p/>`;
    const assertions = [];
    const result = crawl(testTextIndependentNode);
    if (result === undefined) {
        assertions.push("undefined result");
    }
    if (result && result.nodeType !== "INDEPENDENT_NODE_CONFIRMED") {
        assertions.push(`should return INDEPENDENT_NODE_CONFIRMED instead of ${result.nodeType}`);
    }
    if (result && result.vector.origin.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.vector.origin.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.vector.target.arrayIndex !== 0) {
        assertions.push(`should return end arrayIndex as 0`);
    }
    if (result && result.vector.target.stringIndex !== 3) {
        assertions.push(`should return end stringIndex as 3`);
    }
    return assertions;
};
const findOpenParagraphInTextWithArgs = () => {
    const testTextWithArgs = testTextInterpolator `an ${"example"} <p>${"!"}</p>`;
    const assertions = [];
    const result = crawl(testTextWithArgs);
    if (result === undefined) {
        assertions.push("undefined result");
    }
    if (result && result.nodeType !== "OPEN_NODE_CONFIRMED") {
        assertions.push(`should return OPEN_NODE_CONFIRMED instead of ${result.nodeType}`);
    }
    if (result && result.vector.origin.arrayIndex !== 1) {
        assertions.push(`should return start arrayIndex as 1`);
    }
    if (result && result.vector.origin.stringIndex !== 1) {
        assertions.push(`should return start stringIndex as 1`);
    }
    if (result && result.vector.target.arrayIndex !== 1) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.vector.target.stringIndex !== 3) {
        assertions.push(`should return end stringIndex as 3`);
    }
    return assertions;
};
const notFoundInUgglyMessText = () => {
    const testInvalidUgglyMess = testTextInterpolator `an <${"invalid"}p> example${"!"}`;
    const assertions = [];
    const result = crawl(testInvalidUgglyMess);
    if (result === undefined) {
        assertions.push("undefined result");
    }
    if (result && result.nodeType !== "CONTENT_NODE") {
        assertions.push(`should return CONTENT_NODE instead of ${result.nodeType}`);
    }
    if (result && result.vector.origin.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.vector.origin.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.vector.target.arrayIndex !== 2) {
        assertions.push(`should return end arrayIndex as 2`);
    }
    if (result && result.vector.target.stringIndex !== 0) {
        assertions.push(`should return end stringIndex as 0`);
    }
    return assertions;
};
const notFoundInReallyUgglyMessText = () => {
    const testInvalidUgglyMess = testTextInterpolator `an example${"!"}${"?"}`;
    const assertions = [];
    console.log("really ugly mess");
    console.log(testInvalidUgglyMess);
    const result = crawl(testInvalidUgglyMess);
    console.log(result);
    // if (result === undefined) {
    //   assertions.push("undefined result");
    // }
    // if (result && result.nodeType !== "CONTENT_NODE") {
    //   assertions.push(`should return CONTENT_NODE instead of ${result.nodeType}`);
    // }
    // if (result && result.vector.origin.arrayIndex !== 0) {
    //   assertions.push(`should return start arrayIndex as 0`);
    // }
    // if (result && result.vector.origin.stringIndex !== 0) {
    //   assertions.push(`should return start stringIndex as 0`);
    // }
    // if (result && result.vector.target.arrayIndex !== 2) {
    //   assertions.push(`should return end arrayIndex as 2`);
    // }
    // if (result && result.vector.target.stringIndex !== -1) {
    //   assertions.push(`should return end stringIndex as -1`);
    // }
    return assertions;
};
const invalidCloseNodeWithArgs = () => {
    const testInvlaidCloseNodeWithArgs = testTextInterpolator `closed </${"example"}p>`;
    const assertions = [];
    const result = crawl(testInvlaidCloseNodeWithArgs);
    console.log(result);
    if (result === undefined) {
        assertions.push("undefined result");
    }
    if (result && result.nodeType !== "CONTENT_NODE") {
        assertions.push(`should return CONTENT_NODE instead of ${result.nodeType}`);
    }
    if (result && result.vector.origin.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.vector.origin.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.vector.target.arrayIndex !== 1) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.vector.target.stringIndex !== 1) {
        assertions.push(`should return end stringIndex as 1`);
    }
    return assertions;
};
const validCloseNodeWithArgs = () => {
    const testValidCloseNodeWithArgs = testTextInterpolator `closed </p ${"example"}>`;
    const assertions = [];
    const result = crawl(testValidCloseNodeWithArgs);
    console.log(result);
    if (result === undefined) {
        assertions.push("undefined result");
    }
    if (result && result.nodeType !== "CLOSE_NODE_CONFIRMED") {
        assertions.push(`should return CLOSE_NODE_CONFIRMED instead of ${result.nodeType}`);
    }
    if (result && result.vector.origin.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.vector.origin.stringIndex !== 7) {
        assertions.push(`should return start stringIndex as 7`);
    }
    if (result && result.vector.target.arrayIndex !== 1) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.vector.target.stringIndex !== 0) {
        assertions.push(`should return end stringIndex as 0`);
    }
    return assertions;
};
const invalidIndependentNodeWithArgs = () => {
    const testInvalidIndependentNode = testTextInterpolator `independent <${"example"}p/>`;
    const assertions = [];
    const result = crawl(testInvalidIndependentNode);
    console.log(result);
    if (result === undefined) {
        assertions.push("undefined result");
    }
    if (result && result.nodeType !== "CONTENT_NODE") {
        assertions.push(`should return CONTENT_NODE instead of ${result.nodeType}`);
    }
    if (result && result.vector.origin.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.vector.origin.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.vector.target.arrayIndex !== 1) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.vector.target.stringIndex !== 2) {
        assertions.push(`should return end stringIndex as 2`);
    }
    return assertions;
};
const validIndependentNodeWithArgs = () => {
    const testValidIndependentNode = testTextInterpolator `independent <p ${"example"} / >`;
    const assertions = [];
    const result = crawl(testValidIndependentNode);
    console.log(result);
    if (result === undefined) {
        assertions.push("undefined result");
    }
    if (result && result.nodeType !== "INDEPENDENT_NODE_CONFIRMED") {
        assertions.push(`should return INDEPENDENT_NODE_CONFIRMED instead of ${result.nodeType}`);
    }
    if (result && result.vector.origin.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.vector.origin.stringIndex !== 12) {
        assertions.push(`should return start stringIndex as 12`);
    }
    if (result && result.vector.target.arrayIndex !== 1) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.vector.target.stringIndex !== 3) {
        assertions.push(`should return end stringIndex as 3`);
    }
    return assertions;
};
const invalidOpenNodeWithArgs = () => {
    const testInvalidOpenNode = testTextInterpolator `open <${"example"}p>`;
    const assertions = [];
    const result = crawl(testInvalidOpenNode);
    console.log(result);
    if (result === undefined) {
        assertions.push("undefined result");
    }
    if (result && result.nodeType !== "CONTENT_NODE") {
        assertions.push(`should return CONTENT_NODE instead of ${result.nodeType}`);
    }
    if (result && result.vector.origin.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.vector.origin.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.vector.target.arrayIndex !== 1) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.vector.target.stringIndex !== 1) {
        assertions.push(`should return end stringIndex as 1`);
    }
    return assertions;
};
const validOpenNodeWithArgs = () => {
    const testValidOpenNode = testTextInterpolator `open <p ${"example"}>`;
    const assertions = [];
    const result = crawl(testValidOpenNode);
    console.log(result);
    if (result === undefined) {
        assertions.push("undefined result");
    }
    if (result && result.nodeType !== "OPEN_NODE_CONFIRMED") {
        assertions.push(`should return OPEN_NODE_CONFIRMED instead of ${result.nodeType}`);
    }
    if (result && result.vector.origin.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.vector.origin.stringIndex !== 5) {
        assertions.push(`should return start stringIndex as 5`);
    }
    if (result && result.vector.target.arrayIndex !== 1) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.vector.target.stringIndex !== 0) {
        assertions.push(`should return end stringIndex as 0`);
    }
    return assertions;
};
const findNextCrawlWithPreviousCrawl = () => {
    const testValidOpenNode = testTextInterpolator `<p ${"small"}/>${"example"}<p/>`;
    const assertions = [];
    const previousCrawl = crawl(testValidOpenNode);
    const result = crawl(testValidOpenNode, previousCrawl);
    console.log(result);
    if (result === undefined) {
        assertions.push("undefined result");
    }
    if (result && result.nodeType !== "INDEPENDENT_NODE_CONFIRMED") {
        assertions.push(`should return INDEPENDENT_NODE_CONFIRMED instead of ${result.nodeType}`);
    }
    if (result && result.vector.origin.arrayIndex !== 2) {
        assertions.push(`should return start arrayIndex as 2`);
    }
    if (result && result.vector.origin.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.vector.target.arrayIndex !== 2) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.vector.target.stringIndex !== 3) {
        assertions.push(`should return end stringIndex as 3`);
    }
    return assertions;
};
const tests = [
    findNothingWhenThereIsPlainText,
    findParagraphInPlainText,
    findCloseParagraphInPlainText,
    findIndependentParagraphInPlainText,
    findOpenParagraphInTextWithArgs,
    notFoundInUgglyMessText,
    notFoundInReallyUgglyMessText,
    invalidCloseNodeWithArgs,
    validCloseNodeWithArgs,
    invalidIndependentNodeWithArgs,
    validIndependentNodeWithArgs,
    invalidOpenNodeWithArgs,
    validOpenNodeWithArgs,
    findNextCrawlWithPreviousCrawl,
];
const unitTestCrawl = {
    title,
    tests,
    runTestsAsynchronously,
};

// brian taylor vann
const tests$1 = [
    // unitTestRouters,
    unitTestCrawl,
];

// brian taylor vann
const testCollection = [...tests$1];
runTests({ testCollection })
    .then((results) => console.log("results: ", results))
    .catch((errors) => console.log("errors: ", errors));
