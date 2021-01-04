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
const SPECIAL_CHARACTERS = ["_", "-", ".", ":"];
const createAlphaNumericKeys = (route) => {
    const alphabetSet = {};
    const lowercaseLimit = "z".charCodeAt(0);
    const uppercaseLimit = "Z".charCodeAt(0);
    // add letters to seive
    let lowercaseIndex = "a".charCodeAt(0);
    let uppercaseIndex = "A".charCodeAt(0);
    while (lowercaseIndex <= lowercaseLimit) {
        alphabetSet[String.fromCharCode(lowercaseIndex)] = route;
        lowercaseIndex += 1;
    }
    while (uppercaseIndex <= uppercaseLimit) {
        alphabetSet[String.fromCharCode(uppercaseIndex)] = route;
        uppercaseIndex += 1;
    }
    // add numbers
    let numericKey = 0;
    while (numericKey < 10) {
        alphabetSet[numericKey] = route;
        numericKey += 1;
    }
    // add special characters
    for (const specialChar of SPECIAL_CHARACTERS) {
        alphabetSet[specialChar] = route;
    }
    return alphabetSet;
};
const routers = {
    CONTENT_NODE: {
        "<": "OPEN_NODE",
        DEFAULT: "CONTENT_NODE",
    },
    OPEN_NODE: Object.assign(Object.assign({}, createAlphaNumericKeys("OPEN_NODE_VALID")), { "<": "OPEN_NODE", "/": "CLOSE_NODE", DEFAULT: "CONTENT_NODE" }),
    OPEN_NODE_VALID: {
        "<": "OPEN_NODE",
        "/": "INDEPENDENT_NODE_VALID",
        ">": "OPEN_NODE_CONFIRMED",
        DEFAULT: "OPEN_NODE_VALID",
    },
    CLOSE_NODE: Object.assign(Object.assign({}, createAlphaNumericKeys("CLOSE_NODE_VALID")), { "<": "OPEN_NODE", DEFAULT: "CONTENT_NODE" }),
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
const create = (position = DEFAULT_POSITION) => (Object.assign({}, position));
const copy = create;
const increment = (template, position) => {
    // template boundaries
    const templateLength = template.templateArray.length - 1;
    const chunk = template.templateArray[position.arrayIndex];
    if (chunk === undefined) {
        return;
    }
    if (position.arrayIndex >= templateLength &&
        position.stringIndex >= chunk.length - 1) {
        return;
    }
    // cannot % modulo by 0
    if (chunk.length > 0) {
        position.stringIndex += 1;
        position.stringIndex %= chunk.length;
    }
    if (position.stringIndex === 0) {
        position.arrayIndex += 1;
    }
    return position;
};
const decrement = (template, position) => {
    const templateLength = template.templateArray.length - 1;
    if (position.arrayIndex > templateLength) {
        return;
    }
    if (position.arrayIndex <= 0 && position.stringIndex <= 0) {
        return;
    }
    position.stringIndex -= 1;
    if (position.arrayIndex >= 0 && position.stringIndex < 0) {
        position.arrayIndex -= 1;
        const chunk = template.templateArray[position.arrayIndex];
        position.stringIndex = chunk.length - 1;
        // undefined case akin to divide by zero
        if (chunk === "") {
            position.stringIndex = chunk.length;
        }
    }
    return position;
};
const getCharAtPosition = (template, position) => {
    var _a;
    const templateArray = template.templateArray;
    return (_a = templateArray === null || templateArray === void 0 ? void 0 : templateArray[position.arrayIndex]) === null || _a === void 0 ? void 0 : _a[position.stringIndex];
};

const DEFAULT_POSITION$1 = {
    arrayIndex: 0,
    stringIndex: 0,
};
const create$1 = (position = DEFAULT_POSITION$1) => ({
    origin: Object.assign({}, position),
    target: Object.assign({}, position),
});
const createFollowingVector = (template, vector) => {
    const followingVector = copy$1(vector);
    if (increment(template, followingVector.target)) {
        followingVector.origin = copy(followingVector.target);
        return followingVector;
    }
    return;
};
const copy$1 = (vector) => {
    return {
        origin: copy(vector.origin),
        target: copy(vector.target),
    };
};
const incrementOrigin = (template, vector) => {
    if (increment(template, vector.origin)) {
        return vector;
    }
    return;
};
const incrementTarget = (template, vector) => {
    if (increment(template, vector.target)) {
        return vector;
    }
    return;
};
const decrementTarget = (template, vector) => {
    if (decrement(template, vector.target)) {
        return vector;
    }
    return;
};
const hasOriginEclipsedTaraget = (vector) => {
    if (vector.origin.arrayIndex >= vector.target.arrayIndex &&
        vector.origin.stringIndex >= vector.target.stringIndex) {
        return true;
    }
    return false;
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
    let followingVector = create$1();
    if (previousCrawl !== undefined) {
        followingVector = createFollowingVector(template, previousCrawl.vector);
    }
    if (followingVector === undefined) {
        return;
    }
    const crawlState = {
        nodeType: CONTENT_NODE,
        vector: followingVector,
    };
    setNodeType(template, crawlState);
    return crawlState;
};
const setNodeType = (template, crawlState) => {
    var _a, _b;
    const nodeStates = routers[crawlState.nodeType];
    const char = getCharAtPosition(template, crawlState.vector.target);
    if (nodeStates !== undefined && char !== undefined) {
        const defaultNodeType = (_a = nodeStates[DEFAULT]) !== null && _a !== void 0 ? _a : CONTENT_NODE;
        crawlState.nodeType = (_b = nodeStates[char]) !== null && _b !== void 0 ? _b : defaultNodeType;
    }
    return crawlState;
};
const crawl = (template, previousCrawl) => {
    const crawlState = setStartStateProperties(template, previousCrawl);
    if (crawlState === undefined) {
        return;
    }
    let openPosition;
    while (incrementTarget(template, crawlState.vector)) {
        // default to content_node on each cycle
        if (validSieve[crawlState.nodeType] === undefined &&
            crawlState.vector.target.stringIndex === 0) {
            crawlState.nodeType = CONTENT_NODE;
        }
        setNodeType(template, crawlState);
        if (crawlState.nodeType === OPEN_NODE) {
            openPosition = copy(crawlState.vector.target);
        }
        if (confirmedSieve[crawlState.nodeType]) {
            if (openPosition !== undefined) {
                crawlState.vector.origin = openPosition;
            }
            break;
        }
    }
    return crawlState;
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
    const result = crawl(testInvalidUgglyMess);
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

const title$1 = "routers";
const runTestsAsynchronously$1 = true;
const notFoundReducesCorrectState = () => {
    var _a, _b;
    const assertions = [];
    if (((_a = routers["CONTENT_NODE"]) === null || _a === void 0 ? void 0 : _a["<"]) !== "OPEN_NODE") {
        assertions.push("< should return OPEN_NODE");
    }
    if (((_b = routers["CONTENT_NODE"]) === null || _b === void 0 ? void 0 : _b["DEFAULT"]) !== "CONTENT_NODE") {
        assertions.push("space should return CONTENT_NODE");
    }
    return assertions;
};
const openNodeReducesCorrectState = () => {
    var _a, _b, _c, _d;
    const assertions = [];
    if (((_a = routers["OPEN_NODE"]) === null || _a === void 0 ? void 0 : _a["<"]) !== "OPEN_NODE") {
        assertions.push("< should return OPEN_NODE");
    }
    if (((_b = routers["OPEN_NODE"]) === null || _b === void 0 ? void 0 : _b["/"]) !== "CLOSE_NODE") {
        assertions.push("/ should return CLOSE_NODE");
    }
    if (((_c = routers["OPEN_NODE"]) === null || _c === void 0 ? void 0 : _c["b"]) !== "OPEN_NODE_VALID") {
        assertions.push("b should return OPEN_NODE_VALID");
    }
    if (((_d = routers["OPEN_NODE"]) === null || _d === void 0 ? void 0 : _d["DEFAULT"]) !== "CONTENT_NODE") {
        assertions.push("space should return CONTENT_NODE");
    }
    return assertions;
};
const openNodeValidReducesCorrectState = () => {
    var _a, _b, _c, _d;
    const assertions = [];
    if (((_a = routers["OPEN_NODE_VALID"]) === null || _a === void 0 ? void 0 : _a["<"]) !== "OPEN_NODE") {
        assertions.push("< should return OPEN_NODE");
    }
    if (((_b = routers["OPEN_NODE_VALID"]) === null || _b === void 0 ? void 0 : _b["/"]) !== "INDEPENDENT_NODE_VALID") {
        assertions.push("/ should return INDEPENDENT_NODE_VALID");
    }
    if (((_c = routers["OPEN_NODE_VALID"]) === null || _c === void 0 ? void 0 : _c[">"]) !== "OPEN_NODE_CONFIRMED") {
        assertions.push("> should return OPEN_NODE_CONFIRMED");
    }
    if (((_d = routers["OPEN_NODE_VALID"]) === null || _d === void 0 ? void 0 : _d["DEFAULT"]) !== "OPEN_NODE_VALID") {
        assertions.push("space should return OPEN_NODE_VALID");
    }
    return assertions;
};
const independentNodeValidReducesCorrectState = () => {
    var _a, _b;
    const assertions = [];
    if (((_a = routers["INDEPENDENT_NODE_VALID"]) === null || _a === void 0 ? void 0 : _a["<"]) !== "OPEN_NODE") {
        assertions.push("< should return OPEN_NODE");
    }
    if (((_b = routers["INDEPENDENT_NODE_VALID"]) === null || _b === void 0 ? void 0 : _b["DEFAULT"]) !== "INDEPENDENT_NODE_VALID") {
        assertions.push("space should return INDEPENDENT_NODE_VALID");
    }
    return assertions;
};
const closeNodeReducesCorrectState = () => {
    var _a, _b, _c;
    const assertions = [];
    if (((_a = routers["CLOSE_NODE"]) === null || _a === void 0 ? void 0 : _a["<"]) !== "OPEN_NODE") {
        assertions.push("< should return OPEN_NODE");
    }
    if (((_b = routers["CLOSE_NODE"]) === null || _b === void 0 ? void 0 : _b["a"]) !== "CLOSE_NODE_VALID") {
        assertions.push("'a' should return CLOSE_NODE_VALID");
    }
    if (((_c = routers["CLOSE_NODE"]) === null || _c === void 0 ? void 0 : _c["DEFAULT"]) !== "CONTENT_NODE") {
        assertions.push("space should return CLOSE_NODE_VALID");
    }
    return assertions;
};
const closeNodeValidReducesCorrectState = () => {
    var _a, _b, _c;
    const assertions = [];
    if (((_a = routers["CLOSE_NODE_VALID"]) === null || _a === void 0 ? void 0 : _a["<"]) !== "OPEN_NODE") {
        assertions.push("< should return OPEN_NODE");
    }
    if (((_b = routers["CLOSE_NODE_VALID"]) === null || _b === void 0 ? void 0 : _b[">"]) !== "CLOSE_NODE_CONFIRMED") {
        assertions.push("> should return CLOSE_NODE_CONFIRMED");
    }
    if (((_c = routers["CLOSE_NODE_VALID"]) === null || _c === void 0 ? void 0 : _c["DEFAULT"]) !== "CLOSE_NODE_VALID") {
        assertions.push("space should return CLOSE_NODE_VALID");
    }
    return assertions;
};
const tests$1 = [
    notFoundReducesCorrectState,
    openNodeReducesCorrectState,
    openNodeValidReducesCorrectState,
    independentNodeValidReducesCorrectState,
    closeNodeReducesCorrectState,
    closeNodeValidReducesCorrectState,
];
const unitTestCrawlRouters = {
    title: title$1,
    tests: tests$1,
    runTestsAsynchronously: runTestsAsynchronously$1,
};

// brian taylor vann
const MAX_DEPTH = 128;
const DEFAULT_CRAWL_RESULTS = {
    nodeType: "CONTENT_NODE",
    vector: {
        origin: { arrayIndex: 0, stringIndex: 0 },
        target: { arrayIndex: 0, stringIndex: 0 },
    },
};
const SKELETON_SIEVE = {
    ["OPEN_NODE_CONFIRMED"]: "OPEN_NODE",
    ["INDEPENDENT_NODE_CONFIRMED"]: "INDEPENDENT_NODE",
    ["CLOSE_NODE_CONFIRMED"]: "CLOSE_NODE",
    ["CONTENT_NODE"]: "CONTENT_NODE",
};
const buildMissingStringNode = ({ template, currentCrawl, previousCrawl, }) => {
    // get position values
    const originPos = previousCrawl !== undefined
        ? previousCrawl.vector.target
        : DEFAULT_CRAWL_RESULTS.vector.target;
    const targetPos = currentCrawl.vector.origin;
    // justify text vector distance
    const stringDistance = Math.abs(targetPos.stringIndex - originPos.stringIndex);
    const stringArrayDistance = Math.abs(targetPos.arrayIndex - originPos.arrayIndex);
    if (stringDistance + stringArrayDistance < 2) {
        return;
    }
    // copy and correlate position values
    const origin = previousCrawl === undefined
        ? copy(DEFAULT_CRAWL_RESULTS.vector.target)
        : copy(previousCrawl.vector.target);
    const target = copy(currentCrawl.vector.origin);
    decrement(template, target);
    if (previousCrawl !== undefined) {
        increment(template, origin);
    }
    return {
        nodeType: "CONTENT_NODE",
        vector: {
            origin,
            target,
        },
    };
};
const buildSkeleton = (template) => {
    const skeleton = [];
    let depth = 0;
    let previousCrawl;
    let currentCrawl = crawl(template, previousCrawl);
    while (currentCrawl && depth < MAX_DEPTH) {
        // get string in between crawls
        const stringBone = buildMissingStringNode({
            template,
            previousCrawl,
            currentCrawl,
        });
        if (stringBone) {
            skeleton.push(stringBone);
        }
        if (SKELETON_SIEVE[currentCrawl.nodeType]) {
            skeleton.push(currentCrawl);
        }
        previousCrawl = currentCrawl;
        currentCrawl = crawl(template, previousCrawl);
        depth += 1;
    }
    return skeleton;
};

// brian taylor vann
const title$2 = "build_skeleton";
const runTestsAsynchronously$2 = true;
const testTextInterpolator$1 = (templateArray, ...injections) => {
    return { templateArray, injections };
};
const compareSkeletons = (source, target) => {
    for (const sourceKey in source) {
        const node = source[sourceKey];
        const targetNode = target[sourceKey];
        if (targetNode === undefined) {
            return false;
        }
        if (node.nodeType !== targetNode.nodeType) {
            return false;
        }
        if (node.vector.origin.arrayIndex !== targetNode.vector.origin.arrayIndex ||
            node.vector.origin.stringIndex !== targetNode.vector.origin.stringIndex ||
            node.vector.target.arrayIndex !== targetNode.vector.target.arrayIndex ||
            node.vector.target.stringIndex !== targetNode.vector.target.stringIndex) {
            return false;
        }
    }
    return true;
};
const findNothingWhenThereIsPlainText$1 = () => {
    const assertions = [];
    const sourceSkeleton = [
        {
            nodeType: "CONTENT_NODE",
            vector: {
                target: { arrayIndex: 0, stringIndex: 20 },
                origin: { arrayIndex: 0, stringIndex: 0 },
            },
        },
    ];
    const testBlank = testTextInterpolator$1 `no nodes to be found!`;
    const testSkeleton = buildSkeleton(testBlank);
    if (!compareSkeletons(sourceSkeleton, testSkeleton)) {
        assertions.push("skeletons are not equal");
    }
    return assertions;
};
const findParagraphInPlainText$1 = () => {
    const assertions = [];
    const sourceSkeleton = [
        {
            nodeType: "OPEN_NODE_CONFIRMED",
            vector: {
                target: { arrayIndex: 0, stringIndex: 2 },
                origin: { arrayIndex: 0, stringIndex: 0 },
            },
        },
    ];
    const testOpenNode = testTextInterpolator$1 `<p>`;
    const testSkeleton = buildSkeleton(testOpenNode);
    if (!compareSkeletons(sourceSkeleton, testSkeleton)) {
        assertions.push("skeletons are not equal");
    }
    return assertions;
};
const findComplexFromPlainText = () => {
    const assertions = [];
    const sourceSkeleton = [
        {
            nodeType: "CONTENT_NODE",
            vector: {
                target: { arrayIndex: 0, stringIndex: 4 },
                origin: { arrayIndex: 0, stringIndex: 0 },
            },
        },
        {
            nodeType: "OPEN_NODE_CONFIRMED",
            vector: {
                target: { arrayIndex: 0, stringIndex: 7 },
                origin: { arrayIndex: 0, stringIndex: 5 },
            },
        },
        {
            nodeType: "CONTENT_NODE",
            vector: {
                target: { arrayIndex: 0, stringIndex: 12 },
                origin: { arrayIndex: 0, stringIndex: 8 },
            },
        },
        {
            nodeType: "CLOSE_NODE_CONFIRMED",
            vector: {
                target: { arrayIndex: 0, stringIndex: 16 },
                origin: { arrayIndex: 0, stringIndex: 13 },
            },
        },
    ];
    const testComplexNode = testTextInterpolator$1 `hello<p>world</p>`;
    const testSkeleton = buildSkeleton(testComplexNode);
    if (!compareSkeletons(sourceSkeleton, testSkeleton)) {
        assertions.push("skeletons are not equal");
    }
    return assertions;
};
const findCompoundFromPlainText = () => {
    const assertions = [];
    const sourceSkeleton = [
        {
            nodeType: "OPEN_NODE_CONFIRMED",
            vector: {
                target: { arrayIndex: 0, stringIndex: 3 },
                origin: { arrayIndex: 0, stringIndex: 0 },
            },
        },
        {
            nodeType: "CONTENT_NODE",
            vector: {
                target: { arrayIndex: 0, stringIndex: 8 },
                origin: { arrayIndex: 0, stringIndex: 4 },
            },
        },
        {
            nodeType: "CLOSE_NODE_CONFIRMED",
            vector: {
                target: { arrayIndex: 0, stringIndex: 13 },
                origin: { arrayIndex: 0, stringIndex: 9 },
            },
        },
    ];
    const testComplexNode = testTextInterpolator$1 `<h1>hello</h1>`;
    const testSkeleton = buildSkeleton(testComplexNode);
    if (!compareSkeletons(sourceSkeleton, testSkeleton)) {
        assertions.push("skeletons are not equal");
    }
    return assertions;
};
const findBrokenFromPlainText = () => {
    const assertions = [];
    const sourceSkeleton = [
        {
            nodeType: "CONTENT_NODE",
            vector: {
                target: { arrayIndex: 1, stringIndex: 5 },
                origin: { arrayIndex: 0, stringIndex: 0 },
            },
        },
        {
            nodeType: "CLOSE_NODE_CONFIRMED",
            vector: {
                target: { arrayIndex: 1, stringIndex: 10 },
                origin: { arrayIndex: 1, stringIndex: 6 },
            },
        },
        {
            nodeType: "OPEN_NODE_CONFIRMED",
            vector: {
                target: { arrayIndex: 1, stringIndex: 13 },
                origin: { arrayIndex: 1, stringIndex: 11 },
            },
        },
        {
            nodeType: "CONTENT_NODE",
            vector: {
                target: { arrayIndex: 1, stringIndex: 18 },
                origin: { arrayIndex: 1, stringIndex: 14 },
            },
        },
        {
            nodeType: "CLOSE_NODE_CONFIRMED",
            vector: {
                target: { arrayIndex: 1, stringIndex: 22 },
                origin: { arrayIndex: 1, stringIndex: 19 },
            },
        },
    ];
    const testComplexNode = testTextInterpolator$1 `<${"hello"}h2>hey</h2><p>howdy</p>`;
    const testSkeleton = buildSkeleton(testComplexNode);
    if (!compareSkeletons(sourceSkeleton, testSkeleton)) {
        assertions.push("skeletons are not equal");
    }
    return assertions;
};
const tests$2 = [
    findNothingWhenThereIsPlainText$1,
    findParagraphInPlainText$1,
    findComplexFromPlainText,
    findCompoundFromPlainText,
    findBrokenFromPlainText,
];
const unitTestBuildSkeleton = {
    title: title$2,
    tests: tests$2,
    runTestsAsynchronously: runTestsAsynchronously$2,
};

const testTextInterpolator$2 = (templateArray, ...injections) => {
    return { templateArray, injections };
};
const title$3 = "text_position";
const runTestsAsynchronously$3 = true;
const createTextPosition = () => {
    const assertions = [];
    const vector = create();
    if (vector.stringIndex !== 0) {
        assertions.push("text position string index does not match");
    }
    if (vector.arrayIndex !== 0) {
        assertions.push("text position array index does not match");
    }
    return assertions;
};
const createTextPositionFromPosition = () => {
    const assertions = [];
    const prevPosition = {
        stringIndex: 3,
        arrayIndex: 4,
    };
    const vector = create(prevPosition);
    if (vector.stringIndex !== 3) {
        assertions.push("text position string index does not match");
    }
    if (vector.arrayIndex !== 4) {
        assertions.push("text position array index does not match");
    }
    return assertions;
};
const copyTextPosition = () => {
    const assertions = [];
    const position = { arrayIndex: 2, stringIndex: 3 };
    const copiedPosition = copy(position);
    if (position.stringIndex !== copiedPosition.stringIndex) {
        assertions.push("text position string index does not match");
    }
    if (position.arrayIndex !== copiedPosition.arrayIndex) {
        assertions.push("text position array index does not match");
    }
    return assertions;
};
const incrementTextPosition = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$2 `hello`;
    const position = create();
    increment(structureRender, position);
    if (position.stringIndex !== 1) {
        assertions.push("text position string index does not match");
    }
    if (position.arrayIndex !== 0) {
        assertions.push("text position array index does not match");
    }
    return assertions;
};
const incrementMultiTextPosition = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$2 `hey${"world"}, how are you?`;
    const position = create();
    increment(structureRender, position);
    increment(structureRender, position);
    increment(structureRender, position);
    increment(structureRender, position);
    increment(structureRender, position);
    if (position.stringIndex !== 2) {
        assertions.push("text position string index does not match");
    }
    if (position.arrayIndex !== 1) {
        assertions.push("text position array index does not match");
    }
    return assertions;
};
const incrementEmptyTextPosition = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$2 `${"hey"}${"world"}${"!!"}`;
    const position = create();
    increment(structureRender, position);
    increment(structureRender, position);
    increment(structureRender, position);
    if (increment(structureRender, position) !== undefined) {
        assertions.push("should not return after traversed");
    }
    if (position.stringIndex !== 0) {
        assertions.push("text position string index does not match");
    }
    if (position.arrayIndex !== 3) {
        assertions.push("text position array index does not match");
    }
    return assertions;
};
const incrementTextPositionTooFar = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$2 `hey${"world"}, how are you?`;
    const arrayLength = structureRender.templateArray.length - 1;
    const stringLength = structureRender.templateArray[arrayLength].length - 1;
    const position = copy({
        arrayIndex: arrayLength,
        stringIndex: stringLength,
    });
    const MAX_DEPTH = 20;
    let safety = 0;
    while (increment(structureRender, position) && safety < MAX_DEPTH) {
        // iterate across structure
        safety += 1;
    }
    if (position.stringIndex !== 13) {
        assertions.push("text position string index does not match");
    }
    if (position.arrayIndex !== 1) {
        assertions.push("text position array index does not match");
    }
    return assertions;
};
const decrementTextPosition = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$2 `hello`;
    const arrayLength = structureRender.templateArray.length - 1;
    const stringLength = structureRender.templateArray[arrayLength].length - 1;
    const position = copy({
        arrayIndex: arrayLength,
        stringIndex: stringLength,
    });
    decrement(structureRender, position);
    if (position.stringIndex !== 3) {
        assertions.push("text position string index does not match");
    }
    if (position.arrayIndex !== 0) {
        assertions.push("text position array index does not match");
    }
    return assertions;
};
const decrementMultiTextPosition = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$2 `hey${"hello"}bro!`;
    const arrayLength = structureRender.templateArray.length - 1;
    const stringLength = structureRender.templateArray[arrayLength].length - 1;
    const position = copy({
        arrayIndex: arrayLength,
        stringIndex: stringLength,
    });
    decrement(structureRender, position);
    decrement(structureRender, position);
    decrement(structureRender, position);
    decrement(structureRender, position);
    decrement(structureRender, position);
    if (position.stringIndex !== 1) {
        assertions.push("text position string index does not match");
    }
    if (position.arrayIndex !== 0) {
        assertions.push("text position array index does not match");
    }
    return assertions;
};
const decrementEmptyTextPosition = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$2 `${"hey"}${"world"}${"!!"}`;
    const arrayLength = structureRender.templateArray.length - 1;
    const stringLength = structureRender.templateArray[arrayLength].length - 1;
    const position = copy({
        arrayIndex: arrayLength,
        stringIndex: stringLength,
    });
    decrement(structureRender, position);
    decrement(structureRender, position);
    decrement(structureRender, position);
    if (decrement(structureRender, position) !== undefined) {
        assertions.push("should not return after traversed");
    }
    if (position.stringIndex !== 0) {
        assertions.push("text position string index does not match");
    }
    if (position.arrayIndex !== 0) {
        assertions.push("text position array index does not match");
    }
    return assertions;
};
const decrementTextPositionTooFar = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$2 `hey${"world"}, how are you?`;
    const position = create();
    const MAX_DEPTH = 20;
    let safety = 0;
    while (decrement(structureRender, position) && safety < MAX_DEPTH) {
        // iterate across structure
        safety += 1;
    }
    if (position.stringIndex !== 0) {
        assertions.push("text position string index does not match");
    }
    if (position.arrayIndex !== 0) {
        assertions.push("text position array index does not match");
    }
    return assertions;
};
const getCharFromTemplate = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$2 `hello`;
    const position = { arrayIndex: 0, stringIndex: 2 };
    const char = getCharAtPosition(structureRender, position);
    if (char !== "l") {
        assertions.push("textPosition target is not 'l'");
    }
    return assertions;
};
const tests$3 = [
    createTextPosition,
    createTextPositionFromPosition,
    copyTextPosition,
    incrementTextPosition,
    incrementMultiTextPosition,
    incrementEmptyTextPosition,
    incrementTextPositionTooFar,
    decrementTextPosition,
    decrementMultiTextPosition,
    decrementEmptyTextPosition,
    decrementTextPositionTooFar,
    getCharFromTemplate,
];
const unitTestTextPosition = {
    title: title$3,
    tests: tests$3,
    runTestsAsynchronously: runTestsAsynchronously$3,
};

const testTextInterpolator$3 = (templateArray, ...injections) => {
    return { templateArray, injections };
};
const title$4 = "text_vector";
const runTestsAsynchronously$4 = true;
const createTextVector = () => {
    const assertions = [];
    const vector = create$1();
    if (vector.origin.stringIndex !== 0) {
        assertions.push("text vector string index does not match");
    }
    if (vector.origin.arrayIndex !== 0) {
        assertions.push("text vector array index does not match");
    }
    return assertions;
};
const createTextVectorFromPosition = () => {
    const assertions = [];
    const prevPosition = {
        stringIndex: 3,
        arrayIndex: 4,
    };
    const vector = create$1(prevPosition);
    if (vector.origin.stringIndex !== 3) {
        assertions.push("text vector string index does not match");
    }
    if (vector.origin.arrayIndex !== 4) {
        assertions.push("text vector array index does not match");
    }
    return assertions;
};
const copyTextVector = () => {
    const assertions = [];
    const vector = {
        origin: { arrayIndex: 0, stringIndex: 1 },
        target: { arrayIndex: 2, stringIndex: 3 },
    };
    const copiedVector = copy$1(vector);
    if (vector.origin.stringIndex !== copiedVector.origin.stringIndex) {
        assertions.push("text vector string index does not match");
    }
    if (vector.origin.arrayIndex !== copiedVector.origin.arrayIndex) {
        assertions.push("text vector array index does not match");
    }
    return assertions;
};
const incrementTextVector = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$3 `hello`;
    const vector = create$1();
    incrementTarget(structureRender, vector);
    if (vector.target.stringIndex !== 1) {
        assertions.push("text vector string index does not match");
    }
    if (vector.target.arrayIndex !== 0) {
        assertions.push("text vector array index does not match");
    }
    return assertions;
};
const incrementMultiTextVector = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$3 `hey${"world"}, how are you?`;
    const vector = create$1();
    incrementTarget(structureRender, vector);
    incrementTarget(structureRender, vector);
    incrementTarget(structureRender, vector);
    incrementTarget(structureRender, vector);
    incrementTarget(structureRender, vector);
    if (vector.target.stringIndex !== 2) {
        assertions.push("text vector string index does not match");
    }
    if (vector.target.arrayIndex !== 1) {
        assertions.push("text vector array index does not match");
    }
    return assertions;
};
const incrementEmptyTextVector = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$3 `${"hey"}${"world"}${"!!"}`;
    const vector = create$1();
    incrementTarget(structureRender, vector);
    incrementTarget(structureRender, vector);
    incrementTarget(structureRender, vector);
    if (incrementTarget(structureRender, vector) !== undefined) {
        assertions.push("should not return after traversed");
    }
    if (vector.target.stringIndex !== 0) {
        assertions.push("text vector string index does not match");
    }
    if (vector.target.arrayIndex !== 3) {
        assertions.push("text vector array index does not match");
    }
    return assertions;
};
const createFollowingTextVector = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$3 `supercool`;
    const vector = create$1();
    incrementTarget(structureRender, vector);
    incrementTarget(structureRender, vector);
    incrementTarget(structureRender, vector);
    incrementTarget(structureRender, vector);
    const followingVector = createFollowingVector(structureRender, vector);
    if (followingVector === undefined) {
        assertions.push("next vector should not be undefined");
    }
    if (followingVector && followingVector.target.stringIndex !== 5) {
        assertions.push("text vector string index does not match");
    }
    if (followingVector && followingVector.target.arrayIndex !== 0) {
        assertions.push("text vector array index does not match");
    }
    return assertions;
};
const incrementTextVectorTooFar = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$3 `hey${"world"}, how are you?`;
    const vector = create$1();
    const MAX_DEPTH = 20;
    let safety = 0;
    while (incrementTarget(structureRender, vector) && safety < MAX_DEPTH) {
        // iterate across structure
        safety += 1;
    }
    if (vector.target.stringIndex !== 13) {
        assertions.push("text vector string index does not match");
    }
    if (vector.target.arrayIndex !== 1) {
        assertions.push("text vector array index does not match");
    }
    return assertions;
};
const testHasOriginEclipsedTaraget = () => {
    const assertions = [];
    const vector = create$1();
    const results = hasOriginEclipsedTaraget(vector);
    if (results !== true) {
        assertions.push("orign eclipsed target");
    }
    return assertions;
};
const testHasOriginNotEclipsedTaraget = () => {
    const assertions = [];
    const structureRender = testTextInterpolator$3 `hey${"world"}, how are you?`;
    const vector = create$1();
    incrementTarget(structureRender, vector);
    incrementTarget(structureRender, vector);
    incrementTarget(structureRender, vector);
    incrementTarget(structureRender, vector);
    const results = hasOriginEclipsedTaraget(vector);
    if (results !== false) {
        assertions.push("orign has not eclipsed target");
    }
    return assertions;
};
const tests$4 = [
    createTextVector,
    createTextVectorFromPosition,
    createFollowingTextVector,
    copyTextVector,
    incrementTextVector,
    incrementMultiTextVector,
    incrementEmptyTextVector,
    incrementTextVectorTooFar,
    testHasOriginEclipsedTaraget,
    testHasOriginNotEclipsedTaraget,
];
const unitTestTextVector = {
    title: title$4,
    tests: tests$4,
    runTestsAsynchronously: runTestsAsynchronously$4,
};

// brian taylor vann
const crawlForTagName = (template, innerXmlBounds) => {
    const tagVector = copy$1(innerXmlBounds);
    let positionChar = getCharAtPosition(template, tagVector.target);
    if (positionChar === undefined || positionChar === " ") {
        return;
    }
    while (positionChar !== " " && !hasOriginEclipsedTaraget(tagVector)) {
        if (incrementOrigin(template, tagVector) === undefined) {
            return;
        }
        positionChar = getCharAtPosition(template, tagVector.target);
        if (positionChar === undefined) {
            return;
        }
    }
    const adjustedVector = {
        origin: Object.assign({}, innerXmlBounds.origin),
        target: Object.assign({}, tagVector.origin),
    };
    // walk back a step if successive space found
    if (positionChar === " ") {
        decrementTarget(template, adjustedVector);
    }
    return tagVector;
};

// brian taylor vann
const RECURSION_SAFETY = 256;
const testTextInterpolator$4 = (templateArray, ...injections) => {
    return { templateArray, injections };
};
const title$5 = "tag_name_crawl";
const runTestsAsynchronously$5 = true;
const testEmptyString = () => {
    const assertions = [];
    const template = testTextInterpolator$4 ``;
    const vector = create$1();
    const results = crawlForTagName(template, vector);
    if (results !== undefined) {
        assertions.push("this should have failed");
    }
    return assertions;
};
const testEmptySpaceString = () => {
    const assertions = [];
    const template = testTextInterpolator$4 ` `;
    const vector = create$1();
    const results = crawlForTagName(template, vector);
    if (results !== undefined) {
        assertions.push("this should have failed");
    }
    return assertions;
};
const testSingleCharacterString = () => {
    const assertions = [];
    const template = testTextInterpolator$4 `a`;
    const vector = create$1();
    const results = crawlForTagName(template, vector);
    if (results === undefined) {
        assertions.push("this should have returned a vector");
    }
    if (results !== undefined && results.origin.arrayIndex !== 0) {
        assertions.push("incorrect origin array index");
    }
    if (results !== undefined && results.origin.stringIndex !== 0) {
        assertions.push("incorrect origin string index");
    }
    if (results !== undefined && results.target.arrayIndex !== 0) {
        assertions.push("incorrect target array index");
    }
    if (results !== undefined && results.target.stringIndex !== 0) {
        assertions.push("incorrect target string index");
    }
    return assertions;
};
const testCharaceterString = () => {
    const assertions = [];
    const template = testTextInterpolator$4 `a `;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
        safety += 1;
    }
    const results = crawlForTagName(template, vector);
    if (results !== undefined) {
        assertions.push("this should have returned a vector");
    }
    if (results !== undefined && results.origin.arrayIndex !== 0) {
        assertions.push("incorrect origin array index");
    }
    if (results !== undefined && results.origin.stringIndex !== 0) {
        assertions.push("incorrect origin string index");
    }
    if (results !== undefined && results.target.arrayIndex !== 0) {
        assertions.push("incorrect target array index");
    }
    if (results !== undefined && results.target.stringIndex !== 0) {
        assertions.push("incorrect target string index");
    }
    return assertions;
};
const testMultiCharaceterString = () => {
    const assertions = [];
    const template = testTextInterpolator$4 `aaa `;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
        safety += 1;
    }
    const results = crawlForTagName(template, vector);
    if (results !== undefined) {
        assertions.push("this should have returned a vector");
    }
    if (results !== undefined && results.origin.arrayIndex !== 0) {
        assertions.push("incorrect origin array index");
    }
    if (results !== undefined && results.origin.stringIndex !== 0) {
        assertions.push("incorrect origin string index");
    }
    if (results !== undefined && results.target.arrayIndex !== 0) {
        assertions.push("incorrect target array index");
    }
    if (results !== undefined && results.target.stringIndex !== 2) {
        assertions.push("incorrect target string index");
    }
    return assertions;
};
const tests$5 = [
    testEmptyString,
    testEmptySpaceString,
    testSingleCharacterString,
    testCharaceterString,
    testMultiCharaceterString,
];
const unitTestTagNameCrawl = {
    title: title$5,
    tests: tests$5,
    runTestsAsynchronously: runTestsAsynchronously$5,
};

// brian taylor vann
const ATTRIBUTE_FOUND = "ATTRIBUTE_FOUND";
const ATTRIBUTE_ASSIGNMENT = "ATTRIBUTE_ASSIGNMENT";
const CREATE_IMPLICIT_ATTRIBUTE = "CREATE_IMPLICIT_ATTRIBUTE";
const CREATE_EXPLICIT_ATTRIBUTE = "CREATE_EXPLICIT_ATTRIBUTE";
const CREATE_INJECTED_EXPLICIT_ATTRIBUTE = "CREATE_INJECTED_EXPLICIT_ATTRIBUTE";
const getAttributeName = (template, vectorBounds) => {
    const attributeVector = copy$1(vectorBounds);
    let positionChar = getCharAtPosition(template, attributeVector.origin);
    if (positionChar === undefined || positionChar === " ") {
        return;
    }
    let tagNameCrawlState = ATTRIBUTE_FOUND;
    if (positionChar === " ") {
        tagNameCrawlState = CREATE_IMPLICIT_ATTRIBUTE;
    }
    if (positionChar === "=") {
        tagNameCrawlState = ATTRIBUTE_ASSIGNMENT;
    }
    while (tagNameCrawlState === ATTRIBUTE_FOUND &&
        !hasOriginEclipsedTaraget(attributeVector)) {
        if (incrementOrigin(template, attributeVector) === undefined) {
            return;
        }
        positionChar = getCharAtPosition(template, attributeVector.origin);
        if (positionChar === undefined) {
            return;
        }
        tagNameCrawlState = ATTRIBUTE_FOUND;
        if (positionChar === " ") {
            tagNameCrawlState = CREATE_IMPLICIT_ATTRIBUTE;
        }
        if (positionChar === "=") {
            tagNameCrawlState = ATTRIBUTE_ASSIGNMENT;
        }
    }
    // we have found a tag, copy vector
    const adjustedVector = {
        origin: Object.assign({}, vectorBounds.origin),
        target: Object.assign({}, attributeVector.origin),
    };
    if (tagNameCrawlState === ATTRIBUTE_FOUND) {
        return {
            action: CREATE_IMPLICIT_ATTRIBUTE,
            params: { attributeVector: adjustedVector },
        };
    }
    if (tagNameCrawlState === CREATE_IMPLICIT_ATTRIBUTE) {
        if (positionChar === " ") {
            decrementTarget(template, adjustedVector);
        }
        return {
            action: CREATE_IMPLICIT_ATTRIBUTE,
            params: { attributeVector: adjustedVector },
        };
    }
    if (tagNameCrawlState === ATTRIBUTE_ASSIGNMENT) {
        decrementTarget(template, adjustedVector);
        return {
            action: CREATE_EXPLICIT_ATTRIBUTE,
            params: { attributeVector: adjustedVector, valueVector: adjustedVector },
        };
    }
};
const getAttributeQuality = (template, vectorBounds, attributeAction) => {
    // make sure explicity attribute follows (=")
    const attributeVector = copy$1(vectorBounds);
    let positionChar = getCharAtPosition(template, attributeVector.origin);
    if (positionChar !== "=") {
        return;
    }
    incrementOrigin(template, attributeVector);
    if (hasOriginEclipsedTaraget(attributeVector)) {
        return;
    }
    positionChar = getCharAtPosition(template, attributeVector.origin);
    if (positionChar !== '"') {
        return;
    }
    // we have an attribute!
    const attributeQualityVector = copy$1(attributeVector);
    // check for injected attribute
    const arrayIndex = attributeVector.origin.arrayIndex;
    if (incrementOrigin(template, attributeQualityVector) === undefined) {
        return;
    }
    positionChar = getCharAtPosition(template, attributeQualityVector.origin);
    if (positionChar === undefined) {
        return;
    }
    // check if there is a valid injection
    const arrayIndexDistance = Math.abs(arrayIndex - attributeQualityVector.origin.arrayIndex);
    if (arrayIndexDistance > 0 && positionChar !== '"') {
        return;
    }
    if (arrayIndexDistance === 1 && positionChar === '"') {
        // we have an injected attribute
        const injectionVector = {
            origin: Object.assign({}, attributeVector.origin),
            target: Object.assign({}, attributeQualityVector.origin),
        };
        const attributeVectorCopy = copy$1(attributeAction.params.attributeVector);
        return {
            action: CREATE_INJECTED_EXPLICIT_ATTRIBUTE,
            params: {
                attributeVector: attributeVectorCopy,
                valueVector: injectionVector,
                injectionID: arrayIndex,
            },
        };
    }
    // explore potential explicit attribute
    while (positionChar !== '"' &&
        !hasOriginEclipsedTaraget(attributeQualityVector)) {
        if (incrementOrigin(template, attributeQualityVector) === undefined) {
            return;
        }
        // check if valid injection
        if (arrayIndex < attributeQualityVector.origin.arrayIndex) {
            return;
        }
        positionChar = getCharAtPosition(template, attributeQualityVector.origin);
        if (positionChar === undefined) {
            return;
        }
    }
    // check if bounds are valid
    if (positionChar === '"') {
        const explicitVector = {
            origin: Object.assign({}, attributeVector.origin),
            target: Object.assign({}, attributeQualityVector.origin),
        };
        const attributeVectorCopy = copy$1(attributeAction.params.attributeVector);
        return {
            action: "CREATE_EXPLICIT_ATTRIBUTE",
            params: {
                attributeVector: attributeVectorCopy,
                valueVector: explicitVector,
            },
        };
    }
};
const crawlForAttribute = (template, vectorBounds) => {
    // get first character of attribute or return
    const attributeNameResults = getAttributeName(template, vectorBounds);
    if (attributeNameResults === undefined) {
        return;
    }
    if (attributeNameResults.action === "CREATE_IMPLICIT_ATTRIBUTE") {
        return attributeNameResults;
    }
    // get bounding vector
    let qualityVector = copy$1(vectorBounds);
    qualityVector.origin = Object.assign({}, attributeNameResults.params.attributeVector.target);
    incrementOrigin(template, qualityVector);
    return getAttributeQuality(template, qualityVector, attributeNameResults);
};

// brian taylor vann
const RECURSION_SAFETY$1 = 256;
const testTextInterpolator$5 = (templateArray, ...injections) => {
    return { templateArray, injections };
};
const title$6 = "attribute_crawl";
const runTestsAsynchronously$6 = true;
// // but we are seaching between and not incliding '<' '>'
// " " // invalid
//
// // checkbox // invalid
// // checkbox checked // valid
// // checkbox checked  // valid
// // checkbox hello="" // valid
// // checkbox hello="world" // valid
// // checkbox hello="${"world"}" // valid
// // we are looking mainly for ="(-->)"
// //   or we are looking for
const testEmptyString$1 = () => {
    const assertions = [];
    const template = testTextInterpolator$5 ``;
    const vector = create$1();
    const results = crawlForAttribute(template, vector);
    if (results !== undefined) {
        assertions.push("this should have failed");
    }
    return assertions;
};
const testEmptySpaceString$1 = () => {
    const assertions = [];
    const template = testTextInterpolator$5 ` `;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$1) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results !== undefined) {
        assertions.push("this should have failed");
    }
    return assertions;
};
const testEmptyMultiSpaceString = () => {
    const assertions = [];
    const template = testTextInterpolator$5 `   `;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$1) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results !== undefined) {
        assertions.push("this should have failed");
    }
    return assertions;
};
const testImplicitString = () => {
    const assertions = [];
    const template = testTextInterpolator$5 `checked`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$1) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results === undefined) {
        assertions.push("this should not have returned results");
    }
    if (results !== undefined && results.action !== "CREATE_IMPLICIT_ATTRIBUTE") {
        assertions.push("should return CREATE_IMPLICIT_ATTRIBUTE");
    }
    if (results !== undefined && results.action === "CREATE_IMPLICIT_ATTRIBUTE") {
        if (results.params.attributeVector.origin.arrayIndex !== 0) {
            assertions.push("origin.arrayIndex should be 0.");
        }
        if (results.params.attributeVector.origin.stringIndex !== 0) {
            assertions.push("origin.stringIndex should be 0.");
        }
        if (results.params.attributeVector.target.arrayIndex !== 0) {
            assertions.push("target.arrayIndex should be 0.");
        }
        if (results.params.attributeVector.target.stringIndex !== 6) {
            assertions.push("target.stringIndex should be 6.");
        }
    }
    return assertions;
};
const testImplicitStringWithTrailingSpaces = () => {
    const assertions = [];
    const template = testTextInterpolator$5 `checked    `;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$1) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results === undefined) {
        assertions.push("this should not have returned results");
    }
    if (results !== undefined && results.action !== "CREATE_IMPLICIT_ATTRIBUTE") {
        assertions.push("should return CREATE_IMPLICIT_ATTRIBUTE");
    }
    if (results !== undefined && results.action === "CREATE_IMPLICIT_ATTRIBUTE") {
        if (results.params.attributeVector.origin.arrayIndex !== 0) {
            assertions.push("origin.arrayIndex should be 0.");
        }
        if (results.params.attributeVector.origin.stringIndex !== 0) {
            assertions.push("origin.stringIndex should be 0.");
        }
        if (results.params.attributeVector.target.arrayIndex !== 0) {
            assertions.push("target.arrayIndex should be 0.");
        }
        if (results.params.attributeVector.target.stringIndex !== 6) {
            assertions.push("target.stringIndex should be 6.");
        }
    }
    return assertions;
};
const testMalformedExplicitString = () => {
    const assertions = [];
    const template = testTextInterpolator$5 `checked=`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$1) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results !== undefined) {
        assertions.push("this should not have returned results");
    }
    return assertions;
};
const testAlmostExplicitString = () => {
    const assertions = [];
    const template = testTextInterpolator$5 `checked="`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$1) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results !== undefined) {
        assertions.push("this should not have returned results");
    }
    return assertions;
};
const testEmptyExplicitString = () => {
    const assertions = [];
    const template = testTextInterpolator$5 `checked=""`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$1) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results === undefined) {
        assertions.push("this should have returned results");
    }
    if (results !== undefined && results.action === "CREATE_EXPLICIT_ATTRIBUTE") {
        if (results.params.attributeVector.origin.arrayIndex !== 0) {
            assertions.push("attributeVector origin.arrayIndex should be 0.");
        }
        if (results.params.attributeVector.origin.stringIndex !== 0) {
            assertions.push("attributeVector origin.stringIndex should be 0.");
        }
        if (results.params.attributeVector.target.arrayIndex !== 0) {
            assertions.push("attributeVector target.arrayIndex should be 0.");
        }
        if (results.params.attributeVector.target.stringIndex !== 6) {
            assertions.push("attributeVector target.stringIndex should be 6.");
        }
        if (results.params.valueVector.origin.arrayIndex !== 0) {
            assertions.push("valueVector origin.arrayIndex should be 0.");
        }
        if (results.params.valueVector.origin.stringIndex !== 8) {
            assertions.push("valueVector origin.stringIndex should be 0.");
        }
        if (results.params.valueVector.target.arrayIndex !== 0) {
            assertions.push("valueVector target.arrayIndex should be 0.");
        }
        if (results.params.valueVector.target.stringIndex !== 9) {
            assertions.push("valueVector target.stringIndex should be 6.");
        }
    }
    return assertions;
};
const testValidExplicitString = () => {
    const assertions = [];
    const template = testTextInterpolator$5 `checked="checked"`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$1) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results === undefined) {
        assertions.push("this should have returned results");
    }
    if (results !== undefined && results.action === "CREATE_EXPLICIT_ATTRIBUTE") {
        if (results.params.attributeVector.origin.arrayIndex !== 0) {
            assertions.push("attributeVector origin.arrayIndex should be 0.");
        }
        if (results.params.attributeVector.origin.stringIndex !== 0) {
            assertions.push("attributeVector origin.stringIndex should be 0.");
        }
        if (results.params.attributeVector.target.arrayIndex !== 0) {
            assertions.push("attributeVector target.arrayIndex should be 0.");
        }
        if (results.params.attributeVector.target.stringIndex !== 6) {
            assertions.push("attributeVector target.stringIndex should be 6.");
        }
        if (results.params.valueVector.origin.arrayIndex !== 0) {
            assertions.push("valueVector origin.arrayIndex should be 0.");
        }
        if (results.params.valueVector.origin.stringIndex !== 8) {
            assertions.push("valueVector origin.stringIndex should be 0.");
        }
        if (results.params.valueVector.target.arrayIndex !== 0) {
            assertions.push("valueVector target.arrayIndex should be 0.");
        }
        if (results.params.valueVector.target.stringIndex !== 16) {
            assertions.push("valueVector target.stringIndex should be 16.");
        }
    }
    return assertions;
};
const testValidExplicitStringWithTrailingSpaces = () => {
    const assertions = [];
    const template = testTextInterpolator$5 `checked="checked   "`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$1) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results === undefined) {
        assertions.push("this should have returned results");
    }
    if (results !== undefined && results.action === "CREATE_EXPLICIT_ATTRIBUTE") {
        if (results.params.attributeVector.origin.arrayIndex !== 0) {
            assertions.push("attributeVector origin.arrayIndex should be 0.");
        }
        if (results.params.attributeVector.origin.stringIndex !== 0) {
            assertions.push("attributeVector origin.stringIndex should be 0.");
        }
        if (results.params.attributeVector.target.arrayIndex !== 0) {
            assertions.push("attributeVector target.arrayIndex should be 0.");
        }
        if (results.params.attributeVector.target.stringIndex !== 6) {
            assertions.push("attributeVector target.stringIndex should be 6.");
        }
        if (results.params.valueVector.origin.arrayIndex !== 0) {
            assertions.push("valueVector origin.arrayIndex should be 0.");
        }
        if (results.params.valueVector.origin.stringIndex !== 8) {
            assertions.push("valueVector origin.stringIndex should be 0.");
        }
        if (results.params.valueVector.target.arrayIndex !== 0) {
            assertions.push("valueVector target.arrayIndex should be 0.");
        }
        if (results.params.valueVector.target.stringIndex !== 19) {
            assertions.push("valueVector target.stringIndex should be 19.");
        }
    }
    return assertions;
};
const testInjectedString = () => {
    const assertions = [];
    const template = testTextInterpolator$5 `checked="${"hello"}"`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$1) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results === undefined) {
        assertions.push("this should have returned results");
    }
    if (results !== undefined && results.action === "CREATE_EXPLICIT_ATTRIBUTE") {
        if (results.params.attributeVector.origin.arrayIndex !== 0) {
            assertions.push("attributeVector origin.arrayIndex should be 0.");
        }
        if (results.params.attributeVector.origin.stringIndex !== 0) {
            assertions.push("attributeVector origin.stringIndex should be 0.");
        }
        if (results.params.attributeVector.target.arrayIndex !== 0) {
            assertions.push("attributeVector target.arrayIndex should be 0.");
        }
        if (results.params.attributeVector.target.stringIndex !== 6) {
            assertions.push("attributeVector target.stringIndex should be 6.");
        }
        if (results.params.valueVector.origin.arrayIndex !== 0) {
            assertions.push("valueVector origin.arrayIndex should be 0.");
        }
        if (results.params.valueVector.origin.stringIndex !== 8) {
            assertions.push("valueVector origin.stringIndex should be 0.");
        }
        if (results.params.valueVector.target.arrayIndex !== 0) {
            assertions.push("valueVector target.arrayIndex should be 0.");
        }
        if (results.params.valueVector.target.stringIndex !== 19) {
            assertions.push("valueVector target.stringIndex should be 19.");
        }
    }
    return assertions;
};
const testMalformedInjectedString = () => {
    const assertions = [];
    const template = testTextInterpolator$5 `checked="${"hello"}`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$1) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results !== undefined) {
        assertions.push("this should have returned results");
    }
    return assertions;
};
const testMalformedInjectedStringWithTrailingSpaces = () => {
    const assertions = [];
    const template = testTextInterpolator$5 `checked="${"hello"} "`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$1) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results !== undefined) {
        assertions.push("this should not have returned results");
    }
    return assertions;
};
const testMalformedInjectedStringWithStartingSpaces = () => {
    const assertions = [];
    const template = testTextInterpolator$5 `checked=" ${"hello"}"`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$1) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results !== undefined) {
        assertions.push("this should not have returned results");
    }
    return assertions;
};
const tests$6 = [
    testEmptyString$1,
    testEmptySpaceString$1,
    testEmptyMultiSpaceString,
    testImplicitString,
    testImplicitStringWithTrailingSpaces,
    testMalformedExplicitString,
    testAlmostExplicitString,
    testEmptyExplicitString,
    testValidExplicitString,
    testValidExplicitStringWithTrailingSpaces,
    testInjectedString,
    testMalformedInjectedString,
    testMalformedInjectedStringWithTrailingSpaces,
    testMalformedInjectedStringWithStartingSpaces,
];
const unitTestAttributeCrawl = {
    title: title$6,
    tests: tests$6,
    runTestsAsynchronously: runTestsAsynchronously$6,
};

// brian taylor vann
const crawlForContent = (template, vectorBounds) => {
    return {
        action: "CREATE_CONTENT",
        params: { contentVector: copy$1(vectorBounds) },
    };
};

// brian taylor vann
const RECURSION_SAFETY$2 = 256;
const testTextInterpolator$6 = (templateArray, ...injections) => {
    return { templateArray, injections };
};
const title$7 = "content_crawl";
const runTestsAsynchronously$7 = true;
const testEmptyString$2 = () => {
    const assertions = [];
    const template = testTextInterpolator$6 ``;
    const vector = create$1();
    const results = crawlForContent(template, vector);
    if (results === undefined) {
        assertions.push("this should not have failed");
    }
    if (results !== undefined && results.action === "CREATE_CONTENT") {
        if (results.params.contentVector.origin.arrayIndex !== 0) {
            assertions.push("contentVector origin.arrayIndex should be 0.");
        }
        if (results.params.contentVector.origin.stringIndex !== 0) {
            assertions.push("contentVector origin.stringIndex should be 0.");
        }
        if (results.params.contentVector.target.arrayIndex !== 0) {
            assertions.push("contentVector target.arrayIndex should be 0.");
        }
        if (results.params.contentVector.target.stringIndex !== 0) {
            assertions.push("contentVector target.stringIndex should be 0.");
        }
    }
    return assertions;
};
const testLargeEmptyString = () => {
    const assertions = [];
    const template = testTextInterpolator$6 `     `;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$2) {
        safety += 1;
    }
    const results = crawlForContent(template, vector);
    if (results === undefined) {
        assertions.push("this should not have failed");
    }
    if (results !== undefined && results.action === "CREATE_CONTENT") {
        if (results.params.contentVector.origin.arrayIndex !== 0) {
            assertions.push("contentVector origin.arrayIndex should be 0.");
        }
        if (results.params.contentVector.origin.stringIndex !== 0) {
            assertions.push("contentVector origin.stringIndex should be 0.");
        }
        if (results.params.contentVector.target.arrayIndex !== 0) {
            assertions.push("contentVector target.arrayIndex should be 0.");
        }
        if (results.params.contentVector.target.stringIndex !== 4) {
            assertions.push("contentVector target.stringIndex should be 4.");
        }
    }
    return assertions;
};
const testInjectionString = () => {
    const assertions = [];
    const template = testTextInterpolator$6 `${"hello, world!"}`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$2) {
        safety += 1;
    }
    const results = crawlForContent(template, vector);
    if (results === undefined) {
        assertions.push("this should not have failed");
    }
    if (results !== undefined && results.action === "CREATE_CONTENT") {
        if (results.params.contentVector.origin.arrayIndex !== 0) {
            assertions.push("contentVector origin.arrayIndex should be 0.");
        }
        if (results.params.contentVector.origin.stringIndex !== 0) {
            assertions.push("contentVector origin.stringIndex should be 0.");
        }
        if (results.params.contentVector.target.arrayIndex !== 1) {
            assertions.push("contentVector target.arrayIndex should be 1.");
        }
        if (results.params.contentVector.target.stringIndex !== 0) {
            assertions.push("contentVector target.stringIndex should be 0.");
        }
    }
    return assertions;
};
const testLargeInjectionString = () => {
    const assertions = [];
    const template = testTextInterpolator$6 `     ${"hello, world!"}     `;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY$2) {
        safety += 1;
    }
    const results = crawlForContent(template, vector);
    if (results === undefined) {
        assertions.push("this should not have failed");
    }
    if (results !== undefined && results.action === "CREATE_CONTENT") {
        if (results.params.contentVector.origin.arrayIndex !== 0) {
            assertions.push("contentVector origin.arrayIndex should be 0.");
        }
        if (results.params.contentVector.origin.stringIndex !== 0) {
            assertions.push("contentVector origin.stringIndex should be 0.");
        }
        if (results.params.contentVector.target.arrayIndex !== 1) {
            assertions.push("contentVector target.arrayIndex should be 1.");
        }
        if (results.params.contentVector.target.stringIndex !== 4) {
            assertions.push("contentVector target.stringIndex should be 0.");
        }
    }
    return assertions;
};
const tests$7 = [
    testEmptyString$2,
    testLargeEmptyString,
    testInjectionString,
    testLargeInjectionString,
];
const unitTestContentCrawl = {
    title: title$7,
    tests: tests$7,
    runTestsAsynchronously: runTestsAsynchronously$7,
};

// brian taylor vann
const tests$8 = [
    unitTestTextPosition,
    unitTestTextVector,
    unitTestCrawlRouters,
    unitTestCrawl,
    unitTestBuildSkeleton,
    unitTestTagNameCrawl,
    unitTestAttributeCrawl,
    unitTestContentCrawl,
];

// brian taylor vann
const testCollection = [...tests$8];
runTests({ testCollection })
    .then((results) => console.log("results: ", results))
    .catch((errors) => console.log("errors: ", errors));
