import { Subscription, PubSub } from "../../pubsub/pubsub";
import { TestRunResults } from "../state_store/state_types/state_types";

type UnsubscribeToResults = () => void;
type SubscribeToResults = (
  resultsCallback: Subscription<TestRunResults>
) => UnsubscribeToResults;
type BroadcastResults = (testRunState: TestRunResults) => void;

const pubSub = new PubSub<TestRunResults>();

const subscribe: SubscribeToResults = (resultsCallback) => {
  const stub = pubSub.subscribe(resultsCallback);
  return () => {
    pubSub.unsubscribe(stub);
  };
};

// send current state to subscribers
const broadcast: BroadcastResults = (testRunState: TestRunResults) => {
  pubSub.broadcast(testRunState);
};

export { broadcast, subscribe };
