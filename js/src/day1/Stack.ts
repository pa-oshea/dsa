type node<T> = {
    value: T;
    next?: node<T>;
};
export default class Stack<T> {
    public length: number;

    public head?: node<T>;

    constructor() {
        this.head = undefined;
        this.length = 0;
    }

    push(item: T): void {
        const node: node<T> = { value: item };
        this.length++;
        if (!this.head) {
            this.head = node;
            return;
        }
        node.next = this.head;
        this.head = node;
    }

    pop(): T | undefined {
        const result = this.head?.value;

        if (this.head) {
            this.length--;
            this.head = this.head.next;
        }
        return result;
    }

    peek(): T | undefined {
        return this.head?.value;
    }
}
