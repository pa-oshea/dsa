type Node<T> = {
    val: T;
    next?: Node<T>;
};
export default class SinglyLinkedList<T> {
    public length: number;

    public head?: Node<T>;
    public tail?: Node<T>;

    constructor() {
        this.head = this.tail = undefined;
        this.length = 0;
    }

    prepend(item: T): void {
        const node = { val: item } as Node<T>;
        node.next = this.head;
        this.head = node;
        this.length++;
    }

    insertAt(item: T, idx: number): void {
        if (this.length === idx) {
            this.append(item);
            return;
        }

        if (idx === 0) {
            this.prepend(item);
            return;
        }

        let current = this.head;
        while (current && idx > 0) {
            current = current.next;
            idx--;
        }

        if (current) {
            const node: Node<T> = { val: item, next: current.next };
            current.next = node;
            this.length++;
        }
    }

    append(item: T): void {
        const node = { val: item } as Node<T>;
        if (!this.head) {
            this.prepend(item);
            return;
        }
        this.length++;
        if (!this.tail) {
            this.tail = node;
            this.head.next = this.tail;
            return;
        }

        this.tail.next = node;
        this.tail = node;
    }

    remove(item: T): T | undefined {
        if (this.length === 0 || !this.head) {
            return;
        }

        if (this.head.val === item) {
            const result = this.head.val;
            this.head = this.head.next ? this.head.next : undefined;
            this.length--;
            return result;
        }
        let current = this.head;
        while (current.next) {
            if (current.next.val === item) {
                const result = current.next.val;
                current.next = current.next.next
                    ? current.next.next
                    : undefined;
                this.length--;
                return result;
            }
            current = current.next;
        }
        return;
    }

    get(idx: number): T | undefined {
        if (this.length <= idx) {
            return;
        }

        let current = this.head;
        while (current && idx > 0) {
            current = current.next;
            idx--;
        }
        return current?.val;
    }

    removeAt(idx: number): T | undefined {
        const itemToRemove = this.get(idx);
        if (itemToRemove) {
            this.remove(itemToRemove);
        }
        return itemToRemove;
    }
}
