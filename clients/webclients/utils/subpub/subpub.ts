// brian taylor vann

// SubPub
// Dispatch a series of callbacks

type RecycledStubs = Array<number>;
type RecallSubscription<T> = (params: T) => void;
type RecallStore<T> = { [key: string]: RecallSubscription<T> | undefined };

class SubPub<T> {
  private stub = -1;
  private recycledStubs: RecycledStubs = [];
  private subscriptions: RecallStore<T> = {};

  private getStub(): number {
    if (this.recycledStubs.length !== 0) {
      return this.recycledStubs.pop();
    }
    this.stub += 1;
    return this.stub;
  }

  broadcast(params: T): void {
    for (const stubKey in this.subscriptions) {
      const subscription = this.subscriptions[stubKey];
      if (subscription !== undefined) {
        subscription(params);
      }
    }
  }

  subscribe(callback: RecallSubscription<T>): number {
    const stub = this.getStub();
    this.subscriptions[stub] = callback;

    return stub;
  }

  unsubscribe(stub: number): void {
    this.subscriptions[stub] = undefined;
    this.recycledStubs.push(stub);
  }
}

export { RecallSubscription, SubPub };
