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
const ATTRIBUTE_FOUND = "ATTRIBUTE_FOUND";
const IMPLICIT_ATTRIBUTE_CONFIRMED = "IMPLICIT_ATTRIBUTE_CONFIRMED";
const EXPLICIT_ATTRIBUTE_CONFIRMED = "EXPLICIT_ATTRIBUTE_CONFIRMED";
const ATTRIBUTE_ASSIGNMENT = "ATTRIBUTE_ASSIGNMENT";
const getAttributeName = (template, vectorBounds) => {
    const attributeVector = copy$1(vectorBounds);
    let positionChar = getCharAtPosition(template, attributeVector.origin);
    if (positionChar === undefined || positionChar === " ") {
        return;
    }
    let tagNameCrawlState = ATTRIBUTE_FOUND;
    if (positionChar === " ") {
        tagNameCrawlState = IMPLICIT_ATTRIBUTE_CONFIRMED;
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
            tagNameCrawlState = IMPLICIT_ATTRIBUTE_CONFIRMED;
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
            action: IMPLICIT_ATTRIBUTE_CONFIRMED,
            params: { attributeVector: adjustedVector },
        };
    }
    if (tagNameCrawlState === IMPLICIT_ATTRIBUTE_CONFIRMED) {
        if (positionChar === " ") {
            decrementTarget(template, adjustedVector);
        }
        return {
            action: IMPLICIT_ATTRIBUTE_CONFIRMED,
            params: { attributeVector: adjustedVector },
        };
    }
    if (tagNameCrawlState === ATTRIBUTE_ASSIGNMENT) {
        decrementTarget(template, adjustedVector);
        return {
            action: EXPLICIT_ATTRIBUTE_CONFIRMED,
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
            action: "INJECTED_EXPLICIT_ATTRIBUTE_CONFIRMED",
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
            action: "EXPLICIT_ATTRIBUTE_CONFIRMED",
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
    if (attributeNameResults.action === "IMPLICIT_ATTRIBUTE_CONFIRMED") {
        return attributeNameResults;
    }
    // get bounding vector
    let qualityVector = copy$1(vectorBounds);
    qualityVector.origin = Object.assign({}, attributeNameResults.params.attributeVector.target);
    incrementOrigin(template, qualityVector);
    return getAttributeQuality(template, qualityVector, attributeNameResults);
};

// brian taylor vann
const RECURSION_SAFETY = 256;
const testTextInterpolator = (templateArray, ...injections) => {
    return { templateArray, injections };
};
const title = "attribute_crawl";
const runTestsAsynchronously = true;
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
const testEmptyString = () => {
    const assertions = [];
    const template = testTextInterpolator ``;
    const vector = create$1();
    const results = crawlForAttribute(template, vector);
    if (results !== undefined) {
        assertions.push("this should have failed");
    }
    return assertions;
};
const testEmptySpaceString = () => {
    const assertions = [];
    const template = testTextInterpolator ` `;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
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
    const template = testTextInterpolator `   `;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
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
    const template = testTextInterpolator `checked`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results === undefined) {
        assertions.push("this should not have returned results");
    }
    if (results !== undefined &&
        results.action !== "IMPLICIT_ATTRIBUTE_CONFIRMED") {
        assertions.push("should return IMPLICIT_ATTRIBUTE_CONFIRMED");
    }
    if (results !== undefined &&
        results.action === "IMPLICIT_ATTRIBUTE_CONFIRMED") {
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
    const template = testTextInterpolator `checked    `;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results === undefined) {
        assertions.push("this should not have returned results");
    }
    if (results !== undefined &&
        results.action !== "IMPLICIT_ATTRIBUTE_CONFIRMED") {
        assertions.push("should return IMPLICIT_ATTRIBUTE_CONFIRMED");
    }
    if (results !== undefined &&
        results.action === "IMPLICIT_ATTRIBUTE_CONFIRMED") {
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
    const template = testTextInterpolator `checked=`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
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
    const template = testTextInterpolator `checked="`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
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
    const template = testTextInterpolator `checked=""`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results === undefined) {
        assertions.push("this should have returned results");
    }
    if (results !== undefined &&
        results.action === "EXPLICIT_ATTRIBUTE_CONFIRMED") {
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
    const template = testTextInterpolator `checked="checked"`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results === undefined) {
        assertions.push("this should have returned results");
    }
    if (results !== undefined &&
        results.action === "EXPLICIT_ATTRIBUTE_CONFIRMED") {
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
    const template = testTextInterpolator `checked="checked   "`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results === undefined) {
        assertions.push("this should have returned results");
    }
    if (results !== undefined &&
        results.action === "EXPLICIT_ATTRIBUTE_CONFIRMED") {
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
    const template = testTextInterpolator `checked="${"hello"}"`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results === undefined) {
        assertions.push("this should have returned results");
    }
    if (results !== undefined &&
        results.action === "EXPLICIT_ATTRIBUTE_CONFIRMED") {
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
    const template = testTextInterpolator `checked="${"hello"}`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
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
    const template = testTextInterpolator `checked="${"hello"} "`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
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
    const template = testTextInterpolator `checked=" ${"hello"}"`;
    const vector = create$1();
    let safety = 0;
    while (incrementTarget(template, vector) && safety < RECURSION_SAFETY) {
        safety += 1;
    }
    const results = crawlForAttribute(template, vector);
    if (results !== undefined) {
        assertions.push("this should not have returned results");
    }
    return assertions;
};
const tests = [
    testEmptyString,
    testEmptySpaceString,
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
    title,
    tests,
    runTestsAsynchronously,
};

// brian taylor vann
const tests$1 = [
    // unitTestTextPosition,
    // unitTestTextVector,
    // unitTestCrawlRouters,
    // unitTestCrawl,
    // unitTestBuildSkeleton,
    // unitTestTagNameCrawl,
    unitTestAttributeCrawl,
];

// brian taylor vann
const testCollection = [...tests$1];
runTests({ testCollection })
    .then((results) => console.log("results: ", results))
    .catch((errors) => console.log("errors: ", errors));
