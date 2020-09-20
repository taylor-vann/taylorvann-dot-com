// brian taylor vann
// timestamps
let currentTestTimestamp = performance.now();

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

const defaultResultsState = {
    status: "untested",
};
// add test result
// end test result
// create new state based on test collection
const getResults = () => {
    // need to return copy of
    return defaultResultsState;
};

// brian taylor vann
const results = getResults();
console.log("tests!");
console.log("more tests!");
console.log(results);
