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
// needs to be tested
const decrement = (position, template) => {
    const templateLength = template.templateArray.length;
    if (position.arrayIndex < 0 || position.arrayIndex >= templateLength - 1) {
        return;
    }
    const chunkLength = template.templateArray[position.arrayIndex].length;
    if (position.arrayIndex < 0) {
        return;
    }
    position.stringIndex -= 1;
    if (position.stringIndex < 0 && position.arrayIndex > 0) {
        position.arrayIndex -= 1;
        position.stringIndex =
            template.templateArray[position.arrayIndex].length - 1;
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
const crawlHasEnded = (template, previousCrawl) => {
    if (previousCrawl === undefined) {
        return false;
    }
    const templateLength = template.templateArray.length;
    const chunkLength = template.templateArray[previousCrawl.vector.target.arrayIndex].length;
    if (previousCrawl.vector.target.arrayIndex >= templateLength - 1 &&
        previousCrawl.vector.target.stringIndex >= chunkLength - 1) {
        return true;
    }
    return false;
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
    if (crawlHasEnded(template, previousCrawl)) {
        return;
    }
    let openPosition;
    const cState = setStartStateProperties(template, previousCrawl);
    while (increment(cState.vector.target, template)) {
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
const MAX_DEPTH = 128;
const DEFAULT_VECTOR = {
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
const buildMissingStringNode = ({ template, currentCrawl, previousCrawl = DEFAULT_VECTOR, }) => {
    const target = previousCrawl.vector.target;
    const origin = currentCrawl.vector.origin;
    console.log("build missing string node");
    const stringDistance = Math.abs(target.stringIndex - origin.stringIndex);
    const stringArrayDistance = Math.abs(target.arrayIndex - origin.arrayIndex);
    if (stringDistance + stringArrayDistance < 2) {
        console.log("not enough distance to build node");
        return;
    }
    // we need to assess values here
    // copy
    const previousVector = copy(previousCrawl.vector);
    const currentVector = copy(currentCrawl.vector);
    const contentStart = increment(previousVector.target, template);
    const contentEnd = decrement(currentVector.origin, template);
    console.log("do we got content?");
    console.log("build missing string node");
    console.log("start", contentStart);
    console.log("end", contentEnd);
    if (contentStart && contentEnd) {
        console.log("we got content");
        return {
            nodeType: "CONTENT_NODE",
            vector: {
                origin: contentStart,
                target: contentEnd,
            },
        };
    }
};
const buildSkeleton = (template) => {
    let depth = 0;
    const skeleton = [];
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
const title = "build_skeleton";
const runTestsAsynchronously = true;
const testTextInterpolator = (templateArray, ...injections) => {
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
const findNothingWhenThereIsPlainText = () => {
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
    const testBlank = testTextInterpolator `no nodes to be found!`;
    const testSkeleton = buildSkeleton(testBlank);
    if (!compareSkeletons(sourceSkeleton, testSkeleton)) {
        assertions.push("skeletons are not equal");
    }
    return assertions;
};
const findParagraphInPlainText = () => {
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
    const testOpenNode = testTextInterpolator `<p>`;
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
    const testComplexNode = testTextInterpolator `hello<p>world</p>`;
    const testSkeleton = buildSkeleton(testComplexNode);
    console.log(testSkeleton);
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
    const testComplexNode = testTextInterpolator `<h1>hello</h1>`;
    const testSkeleton = buildSkeleton(testComplexNode);
    console.log(testSkeleton);
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
    const testComplexNode = testTextInterpolator `<${"hello"}h2>hey</h2><p>howdy</p>`;
    const testSkeleton = buildSkeleton(testComplexNode);
    console.log(testSkeleton);
    if (!compareSkeletons(sourceSkeleton, testSkeleton)) {
        assertions.push("skeletons are not equal");
    }
    return assertions;
};
const tests = [
    findNothingWhenThereIsPlainText,
    findParagraphInPlainText,
    findComplexFromPlainText,
    findCompoundFromPlainText,
    findBrokenFromPlainText,
];
const unitTestBuildSkeleton = {
    title,
    tests,
    runTestsAsynchronously,
};

// brian taylor vann
const tests$1 = [
    // unitTestRouters,
    // unitTestCrawl,
    unitTestBuildSkeleton,
];

// brian taylor vann
const testCollection = [...tests$1];
runTests({ testCollection })
    .then((results) => console.log("results: ", results))
    .catch((errors) => console.log("errors: ", errors));
