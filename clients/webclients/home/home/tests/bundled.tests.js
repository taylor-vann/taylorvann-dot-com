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

const DEFAULT_POSITION = {
    arrayIndex: 0,
    stringIndex: 0,
};
const create = (position = DEFAULT_POSITION) => (Object.assign({}, position));
const copy = create;
const increment = (template, position) => {
    // template boundaries
    const templateLength = template.templateArray.length;
    const chunkLength = template.templateArray[position.arrayIndex].length;
    if (chunkLength === undefined) {
        return;
    }
    const arrayIndex = position.arrayIndex;
    const stringIndex = position.stringIndex;
    if (arrayIndex >= templateLength - 1 && chunkLength === 0) {
        return;
    }
    if (arrayIndex >= templateLength - 1 && stringIndex >= chunkLength - 1) {
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
// needs to be tested
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
        // base case akin to divide by zero
        if (chunk === "") {
            position.stringIndex = chunk.length;
        }
    }
    return position;
};
const getCharFromTarget = (template, position) => {
    const templateArray = template.templateArray;
    const arrayIndex = position.arrayIndex;
    const stringIndex = position.stringIndex;
    if (0 <= arrayIndex && arrayIndex > templateArray.length - 1) {
        return;
    }
    if (0 <= stringIndex && stringIndex > templateArray[arrayIndex].length - 1) {
        return;
    }
    return templateArray[arrayIndex][stringIndex];
};

const testTextInterpolator = (templateArray, ...injections) => {
    return { templateArray, injections };
};
const title = "text_position";
const runTestsAsynchronously = true;
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
    const structureRender = testTextInterpolator `hello`;
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
    const structureRender = testTextInterpolator `hey${"world"}, how are you?`;
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
    const structureRender = testTextInterpolator `${"hey"}${"world"}${"!!"}`;
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
    const structureRender = testTextInterpolator `hey${"world"}, how are you?`;
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
    const structureRender = testTextInterpolator `hello`;
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
    const structureRender = testTextInterpolator `hey${"hello"}bro!`;
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
    const structureRender = testTextInterpolator `${"hey"}${"world"}${"!!"}`;
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
    const structureRender = testTextInterpolator `hey${"world"}, how are you?`;
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
    const structureRender = testTextInterpolator `hello`;
    const position = { arrayIndex: 0, stringIndex: 2 };
    const char = getCharFromTarget(structureRender, position);
    if (char !== "l") {
        assertions.push("textPosition target is not 'l'");
    }
    return assertions;
};
const tests = [
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
    title,
    tests,
    runTestsAsynchronously,
};

// brian taylor vann
const tests$1 = [
    unitTestTextPosition,
];

// brian taylor vann
const testCollection = [...tests$1];
runTests({ testCollection })
    .then((results) => console.log("results: ", results))
    .catch((errors) => console.log("errors: ", errors));
