type Node<T> = {
    val: T;
    next?: Node<T>;
    prev?: Node<T>;
};
export default class DoublyLinkedList<T> {
    public length: number;
    private head?: Node<T>;
    private tail?: Node<T>;

    constructor() {
        this.head = this.tail = undefined;
        this.length = 0;
    }

    prepend(item: T): void {
        const node: Node<T> = { val: item };
        if (this.length === 0) {
            this.tail = node;
        } else {
            this.head!.prev = node;
            node.next = this.head;
        }
        this.head = node;
        this.length++;
    }

    insertAt(item: T, idx: number): void {
        if (idx > this.length) {
            throw new Error("index out of range");
        } else if (idx === this.length) {
            this.append(item);
            return;
        } else if (idx === 0) {
            this.prepend(item);
            return;
        }

        let curr = this.getAt(idx);

        const node: Node<T> = { val: item, next: curr, prev: curr!.prev };

        curr!.prev!.next = node;
        curr!.prev = node;
        this.length++;
    }

    append(item: T): void {
        const node: Node<T> = { val: item };
        if (this.length === 0) {
            this.head = node;
        } else {
            this.tail!.next = node;
            node.prev = this.tail;
        }

        this.tail = node;
        this.length++;
    }

    remove(item: T): T | undefined {
        let curr = this.head;
        for (let i = 0; curr && curr.val !== item && i < this.length; i++) {
            curr = curr.next;
        }

        if (!curr) {
            return;
        }

        return this.removeNode(curr);
    }

    get(idx: number): T | undefined {
        if (this.length === 0 && idx >= this.length) {
            return;
        }

        // let curr = this.head;
        // let ishead = true;
        // if (idx > this.length) {
        //     curr = this.tail;
        //     ishead = false;
        // }
        let curr = this.getAt(idx);

        return curr?.val;
    }

    removeAt(idx: number): T | undefined {
        const node = this.getAt(idx);
        return node ? this.removeNode(node) : undefined;
    }

    private getAt(idx: number): Node<T> | undefined {
        let curr = this.head;
        // for (let i = 0; curr && i < this.length; i++) {
        //     curr = curr.next;
        // }
        //
        while (curr && idx > 0) {
            curr = curr.next;
            idx--;
        }

        return curr;
    }

    private removeNode(node: Node<T>): T | undefined {
        this.length--;

        if (this.length === 0) {
            const result = this.head?.val;
            this.head = this.tail = undefined;
            return result;
        }
        // if (node.prev) node.prev = node.next;
        //
        // if (node.next) node.next = node.next;
        // node.next = node?.prev;
        // node.prev = node.next;
        if (node === this.head) {
            const result = this.head?.val;
            this.head = node.next;
            return result;
        }

        if (node === this.tail) {
            const result = this.tail?.val;
            this.tail = node.prev;
            return result;
        }
        node.prev!.next = node.next;
        node.next!.prev = node.prev;
        node.next = node.prev = undefined;
        return node.val;
    }
}
