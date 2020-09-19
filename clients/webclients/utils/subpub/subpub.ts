// brian taylor vann

// SubPub
// Dispatch a series of callbacks

type RecycledStubs = Array<number>;
type Subscription<T> = (params: T) => void;
type SubsctiptionStore<T> = { [key: string]: Subscription<T> | undefined };

class SubPub<T> {
  private stub = 0;
  private recycledStubs: RecycledStubs = [];
  private subscriptions: SubsctiptionStore<T> = {};

  private getStub(): number {
    const stub = this.recycledStubs.pop();
    if (stub) {
      return stub;
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

  subscribe(callback: Subscription<T>): number {
    const stub = this.getStub();
    this.subscriptions[stub] = callback;

    return stub;
  }

  unsubscribe(stub: number): void {
    this.subscriptions[stub] = undefined;
    this.recycledStubs.push(stub);
  }
}

export { SubPub };
