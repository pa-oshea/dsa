type node<T> = {
    value: T;
    next?: node<T>;
};

export default class Queue<T> {
    public length: number;

    private head?: node<T>;
    private tail?: node<T>;

    constructor() {
        this.head = this.tail = undefined;
        this.length = 0;
    }

    enqueue(item: T): void {
        const node: node<T> = { value: item };
        this.length++;
        if (!this.tail) {
            this.tail = node;
            this.tail.next = node;
            this.head = this.tail;
            return;
        }
        this.tail.next = node;
        this.tail = node;
    }

    deque(): T | undefined {
        if (!this.head) {
            return;
        }
        const result = this.head?.value;
        this.head = this.head.next;
        this.length--;
        if (this.length === 0) {
            this.tail = undefined;
        }
        return result;
    }

    peek(): T | undefined {
        return this.tail?.value;
    }
}
