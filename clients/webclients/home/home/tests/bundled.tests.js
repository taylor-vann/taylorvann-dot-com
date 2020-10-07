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

const buildResultsState = ({ startTime, testCollection, }) => {
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
let currentTestTimestamp = performance.now();
const getTimestamp = () => {
    return currentTestTimestamp;
};
const updateTimestamp = () => {
    currentTestTimestamp = performance.now();
    return currentTestTimestamp;
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
        if (issuedAt < getTimestamp()) {
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
        if (issuedAt < getTimestamp()) {
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
    if (startTime < getTimestamp()) {
        return;
    }
    yield Promise.all(builtAsyncTests);
});
const runTestsInOrder = ({ startTime, collectionID, tests, timeoutInterval, }) => __awaiter(void 0, void 0, void 0, function* () {
    let testID = 0;
    for (const testFunc of tests) {
        if (startTime < getTimestamp()) {
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
const startLtrTestCollectionRun = ({ testCollection, startTime, }) => __awaiter$1(void 0, void 0, void 0, function* () {
    startTestRun({ testCollection, startTime });
    let collectionID = 0;
    for (const collection of testCollection) {
        if (startTime < getTimestamp()) {
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
        if (startTime < getTimestamp()) {
            return;
        }
        const endTime = performance.now();
        endTestCollection$1({
            collectionID,
            endTime,
        });
        collectionID += 1;
    }
    if (startTime < getTimestamp()) {
        return;
    }
    const endTime = performance.now();
    endTestRun$1({ endTime });
});
// iterate through tests synchronously
const runTests = (params) => __awaiter$1(void 0, void 0, void 0, function* () {
    const startTime = updateTimestamp();
    yield startLtrTestCollectionRun(Object.assign(Object.assign({}, params), { startTime }));
    if (startTime < getTimestamp()) {
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

// brian taylor vann
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
const createNotFoundCrawlState = () => {
    return {
        nodeType: "CONTENT_NODE",
        target: {
            start: {
                arrayIndex: 0,
                stringIndex: 0,
            },
            end: {
                arrayIndex: 0,
                stringIndex: 0,
            },
        },
    };
};
const setStartStateProperties = (brokenText, previousCrawl) => {
    const cState = createNotFoundCrawlState();
    if (previousCrawl === undefined) {
        return cState;
    }
    let { arrayIndex, stringIndex } = previousCrawl.target.end;
    stringIndex += 1;
    stringIndex %= brokenText[arrayIndex].length;
    if (stringIndex === 0) {
        arrayIndex += 1;
    }
    if (arrayIndex >= brokenText.length) {
        return;
    }
    cState.target.start.arrayIndex = arrayIndex;
    cState.target.start.stringIndex = stringIndex;
    cState.target.end.arrayIndex = arrayIndex;
    cState.target.end.stringIndex = stringIndex;
    return cState;
};
const setNodeType = (cState, char) => {
    var _a, _b, _c, _d;
    const defaultNodeType = (_b = (_a = routers[cState.nodeType]) === null || _a === void 0 ? void 0 : _a["DEFAULT"]) !== null && _b !== void 0 ? _b : CONTENT_NODE;
    cState.nodeType = (_d = (_c = routers[cState.nodeType]) === null || _c === void 0 ? void 0 : _c[char]) !== null && _d !== void 0 ? _d : defaultNodeType;
    return cState;
};
const setStart = (results, arrayIndex, stringIndex) => {
    results.target.start.arrayIndex = arrayIndex;
    results.target.start.stringIndex = stringIndex;
    results.target.end.arrayIndex = arrayIndex;
    results.target.end.stringIndex = stringIndex;
};
const setEnd = (results, arrayIndex, stringIndex) => {
    results.target.end.arrayIndex = arrayIndex;
    results.target.end.stringIndex = stringIndex;
};
const crawl = (brokenText, previousCrawl) => {
    const cState = setStartStateProperties(brokenText, previousCrawl);
    if (cState === undefined) {
        return;
    }
    let { stringIndex, arrayIndex } = cState.target.start;
    // retain most recent postition
    const suspect = {
        arrayIndex,
        stringIndex,
    };
    while (arrayIndex < brokenText.length) {
        if (validSieve[cState.nodeType] === undefined) {
            cState.nodeType = CONTENT_NODE;
        }
        const chunk = brokenText[arrayIndex];
        while (stringIndex < chunk.length) {
            setNodeType(cState, chunk.charAt(stringIndex));
            if (confirmedSieve[cState.nodeType]) {
                // if confirmed, suspected target is verified
                setStart(cState, suspect.arrayIndex, suspect.stringIndex);
                setEnd(cState, arrayIndex, stringIndex);
                return cState;
            }
            if (cState.nodeType === OPEN_NODE) {
                suspect.arrayIndex = arrayIndex;
                suspect.stringIndex = stringIndex;
            }
            stringIndex += 1;
        }
        // skip to next chunk
        stringIndex = 0;
        arrayIndex += 1;
    }
    // finished walk without results
    arrayIndex = brokenText.length - 1;
    stringIndex = brokenText[arrayIndex].length - 1;
    setEnd(cState, arrayIndex, stringIndex);
    return cState;
};

// brian taylor vann
const testTextInterpolator = (brokenText, ...injections) => {
    return brokenText;
};
const title = "Crawl";
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
    if (result && result.target.start.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.target.start.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.target.end.arrayIndex !== 0) {
        assertions.push(`should return end arrayIndex as 0`);
    }
    if (result && result.target.end.stringIndex !== 20) {
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
    if (result && result.target.start.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.target.start.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.target.end.arrayIndex !== 0) {
        assertions.push(`should return end arrayIndex as 0`);
    }
    if (result && result.target.end.stringIndex !== 2) {
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
    if (result && result.target.start.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.target.start.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 2`);
    }
    if (result && result.target.end.arrayIndex !== 0) {
        assertions.push(`should return end arrayIndex as 0`);
    }
    if (result && result.target.end.stringIndex !== 3) {
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
    if (result && result.target.start.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.target.start.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.target.end.arrayIndex !== 0) {
        assertions.push(`should return end arrayIndex as 0`);
    }
    if (result && result.target.end.stringIndex !== 3) {
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
    if (result && result.target.start.arrayIndex !== 1) {
        assertions.push(`should return start arrayIndex as 1`);
    }
    if (result && result.target.start.stringIndex !== 1) {
        assertions.push(`should return start stringIndex as 1`);
    }
    if (result && result.target.end.arrayIndex !== 1) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.target.end.stringIndex !== 3) {
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
    if (result && result.target.start.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.target.start.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.target.end.arrayIndex !== 2) {
        assertions.push(`should return end arrayIndex as 2`);
    }
    if (result && result.target.end.stringIndex !== -1) {
        assertions.push(`should return end stringIndex as -1`);
    }
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
    if (result && result.target.start.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.target.start.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.target.end.arrayIndex !== 1) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.target.end.stringIndex !== 1) {
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
    if (result && result.target.start.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.target.start.stringIndex !== 7) {
        assertions.push(`should return start stringIndex as 7`);
    }
    if (result && result.target.end.arrayIndex !== 1) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.target.end.stringIndex !== 0) {
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
    if (result && result.target.start.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.target.start.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.target.end.arrayIndex !== 1) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.target.end.stringIndex !== 2) {
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
    if (result && result.target.start.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.target.start.stringIndex !== 12) {
        assertions.push(`should return start stringIndex as 12`);
    }
    if (result && result.target.end.arrayIndex !== 1) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.target.end.stringIndex !== 3) {
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
    if (result && result.target.start.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.target.start.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.target.end.arrayIndex !== 1) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.target.end.stringIndex !== 1) {
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
    if (result && result.target.start.arrayIndex !== 0) {
        assertions.push(`should return start arrayIndex as 0`);
    }
    if (result && result.target.start.stringIndex !== 5) {
        assertions.push(`should return start stringIndex as 5`);
    }
    if (result && result.target.end.arrayIndex !== 1) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.target.end.stringIndex !== 0) {
        assertions.push(`should return end stringIndex as 0`);
    }
    return assertions;
};
const validSecondaryIndependentNodeWithArgs = () => {
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
    if (result && result.target.start.arrayIndex !== 2) {
        assertions.push(`should return start arrayIndex as 2`);
    }
    if (result && result.target.start.stringIndex !== 0) {
        assertions.push(`should return start stringIndex as 0`);
    }
    if (result && result.target.end.arrayIndex !== 2) {
        assertions.push(`should return end arrayIndex as 1`);
    }
    if (result && result.target.end.stringIndex !== 3) {
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
    invalidCloseNodeWithArgs,
    validCloseNodeWithArgs,
    invalidIndependentNodeWithArgs,
    validIndependentNodeWithArgs,
    invalidOpenNodeWithArgs,
    validOpenNodeWithArgs,
    validSecondaryIndependentNodeWithArgs,
];
const unitTestCrawl = {
    title,
    tests,
    runTestsAsynchronously,
};

const title$1 = "Routers | Detect node state";
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
const unitTestRouters = {
    title: title$1,
    tests: tests$1,
    runTestsAsynchronously: runTestsAsynchronously$1,
};

const MAX_RECURSION = 128;
const SKELETON_SIEVE = {
    ["OPEN_NODE_CONFIRMED"]: "OPEN_NODE",
    ["INDEPENDENT_NODE_CONFIRMED"]: "INDEPENDENT_NODE",
    ["CLOSE_NODE_CONFIRMED"]: "CLOSE_NODE",
    ["CONTENT_NODE"]: "CONTENT_NODE",
};
const getStringBoneStart = (brokenText, previousCrawl) => {
    let { arrayIndex, stringIndex } = previousCrawl.target.end;
    stringIndex += 1;
    stringIndex %= brokenText[arrayIndex].length;
    if (stringIndex === 0) {
        arrayIndex += 1;
    }
    return {
        arrayIndex,
        stringIndex,
    };
};
const getStringBoneEnd = (brokenText, currentCrawl) => {
    let { arrayIndex, stringIndex } = currentCrawl.target.start;
    stringIndex -= 1;
    if (stringIndex === -1) {
        arrayIndex -= 1;
        stringIndex += brokenText[arrayIndex].length;
    }
    return {
        arrayIndex,
        stringIndex,
    };
};
const buildSkeletonStringBone = ({ brokenText, currentCrawl, previousCrawl, }) => {
    if (previousCrawl === undefined) {
        return;
    }
    const { end } = previousCrawl.target;
    const { start } = currentCrawl.target;
    const stringDistance = Math.abs(start.stringIndex - end.stringIndex);
    const stringArrayDistance = start.arrayIndex - end.arrayIndex;
    if (2 > stringArrayDistance + stringDistance) {
        return;
    }
    const contentStart = getStringBoneStart(brokenText, previousCrawl);
    const contentEnd = getStringBoneEnd(brokenText, currentCrawl);
    if (contentStart && contentEnd) {
        return {
            nodeType: "CONTENT_NODE",
            target: {
                start: contentStart,
                end: contentEnd,
            },
        };
    }
};
const buildSkeleton = (brokenText, ...injections) => {
    let depth = 0;
    const skeleton = [];
    let previousCrawl;
    let currentCrawl = crawl(brokenText, previousCrawl);
    while (currentCrawl && depth < MAX_RECURSION) {
        // get string in between crawls
        const stringBone = buildSkeletonStringBone({
            brokenText,
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
        currentCrawl = crawl(brokenText, previousCrawl);
        depth += 1;
    }
    return skeleton;
};

// brian taylor vann
// order of start, end aren't being respected
const title$2 = "bang/xml_crawler/crawl";
const runTestsAsynchronously$2 = true;
// const findNothingWhenThereIsPlainText = () => {
//   const testBlank = buildSkeleton`no nodes to be found!`;
//   console.log(testBlank);
//   const assertions: string[] = [];
//   return assertions;
// };
// const findParagraphInPlainText = () => {
//   const testOpenNode = buildSkeleton`<p>`;
//   console.log(testOpenNode);
//   const assertions: string[] = [];
//   return assertions;
// };
// const findComplexFromPlainText = () => {
//   const testComplexNode = buildSkeleton`hello<p>world</p>`;
//   console.log(testComplexNode);
//   const assertions: string[] = [];
//   return assertions;
// };
// const findCompoundFromPlainText = () => {
//   const testComplexNode = buildSkeleton`<h1>hello</h1><h2>world</h2><img/><p>howdy</p>`;
//   console.log(testComplexNode);
//   const assertions: string[] = [];
//   return assertions;
// };
const findBrokenFromPlainText = () => {
    const testComplexNode = buildSkeleton `<h1>hello</h1><${"hello"}h2>world</h2><p>howdy</p>`;
    console.log(testComplexNode);
    const assertions = [];
    return assertions;
};
// const findCloseParagraphInPlainText = () => {
//   const testTextCloseNode = testTextInterpolator`</p>`;
//   const assertions: string[] = [];
//   return assertions;
// };
// const findIndependentParagraphInPlainText = () => {
//   const testTextIndependentNode = testTextInterpolator`<p/>`;
//   const assertions: string[] = [];
//   return assertions;
// };
// const findOpenParagraphInTextWithArgs = () => {
//   const testTextWithArgs = testTextInterpolator`an ${"example"} <p>${"!"}</p>`;
//   const assertions: string[] = [];
//   return assertions;
// };
// const notFoundInUgglyMessText = () => {
//   const testInvalidUgglyMess = testTextInterpolator`an <${"invalid"}p> example${"!"}`;
//   const assertions: string[] = [];
//   return assertions;
// };
// const invalidCloseNodeWithArgs = () => {
//   const testInvlaidCloseNodeWithArgs = testTextInterpolator`closed </${"example"}p>`;
//   const assertions: string[] = [];
//   return assertions;
// };
// const validCloseNodeWithArgs = () => {
//   const testValidCloseNodeWithArgs = testTextInterpolator`closed </p ${"example"}>`;
//   const assertions: string[] = [];
//   return assertions;
// };
// const invalidIndependentNodeWithArgs = () => {
//   const testInvalidIndependentNode = testTextInterpolator`independent <${"example"}p/>`;
//   const assertions: string[] = [];
//   return assertions;
// };
// const validIndependentNodeWithArgs = () => {
//   const testValidIndependentNode = testTextInterpolator`independent <p ${"example"} / >`;
//   const assertions: string[] = [];
//   return assertions;
// };
// const invalidOpenNodeWithArgs = () => {
//   const testInvalidOpenNode = testTextInterpolator`open <${"example"}p>`;
//   const assertions: string[] = [];
//   return assertions;
// };
// const validOpenNodeWithArgs = () => {
//   const testValidOpenNode = testTextInterpolator`open <p ${"example"}>`;
//   const assertions: string[] = [];
//   return assertions;
// };
const tests$2 = [
    // findNothingWhenThereIsPlainText,
    // findParagraphInPlainText,
    // findComplexFromPlainText,
    // findCompoundFromPlainText,
    findBrokenFromPlainText,
];
const unitTestBuildSkeleton = {
    title: title$2,
    tests: tests$2,
    runTestsAsynchronously: runTestsAsynchronously$2,
};

// brian taylor vann
// import { unitTestXMLCrawler } from "./xml_crawler/xml_crawler.test";
const tests$3 = [unitTestRouters, unitTestCrawl, unitTestBuildSkeleton];

// brian taylor vann
const testCollection = [...tests$3];
runTests({ testCollection })
    .then((results) => console.log("results: ", results))
    .catch((errors) => console.log("errors: ", errors));
