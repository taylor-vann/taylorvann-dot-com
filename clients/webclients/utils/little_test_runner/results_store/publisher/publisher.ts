import { Subscription, SubPub } from "../../subpub/subpub";
import { TestRunResults } from "../state_store/state_store";

type UnsubscribeToResults = () => void;
type SubscribeToResults = (
  resultsCallback: Subscription<TestRunResults>
) => UnsubscribeToResults;
type BroadcastResults = (testRunState: TestRunResults) => void;

const subpub = new SubPub<TestRunResults>();

const subscribe: SubscribeToResults = (resultsCallback) => {
  const stub = subpub.subscribe(resultsCallback);
  return () => {
    subpub.unsubscribe(stub);
  };
};

// send current state to subscribers
const broadcast: BroadcastResults = (testRunState: TestRunResults) => {
  subpub.broadcast(testRunState);
};

export { broadcast, subscribe };
