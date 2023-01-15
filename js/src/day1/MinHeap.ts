export default class MinHeap {
    public length: number;
    private data: number[];

    constructor() {
        this.length = 0;
        this.data = [];
    }

    insert(value: number): void {
        this.data[this.length] = value;
        this.heapifyUp(this.length);
        this.length++;
    }
    delete(): number | undefined {
        if (this.length === 0) {
            return;
        }

        this.length--;
        const out = this.data[0];
        if (this.length === 0) {
            this.data = [];
            return out;
        }
        this.data[0] = this.data[this.length];
        this.heapifyDown(0);

        return out;
    }

    private heapifyDown(idx: number) {
        if (idx >= this.length) {
            return;
        }

        const leftIdx = this.getLeftIdx(idx);
        const rightIdx = this.getRightIdx(idx);

        if (leftIdx >= this.length) {
            return;
        }

        const leftValue = this.data[leftIdx];
        const rightValue = this.data[rightIdx];
        const value = this.data[idx];

        if (leftValue > rightValue && value > rightValue) {
            this.data[idx] = rightValue;
            this.data[rightIdx] = value;
            this.heapifyDown(rightIdx);
        } else if (rightValue > leftValue && value > leftValue) {
            this.data[idx] = leftValue;
            this.data[leftIdx] = value;
            this.heapifyDown(leftIdx);
        }
    }

    private heapifyUp(idx: number) {
        if (idx === 0) {
            return;
        }

        const p = this.getParentIdx(idx);
        const parentValue = this.data[p];
        const value = this.data[idx];

        if (parentValue > value) {
            this.data[p] = value;
            this.data[idx] = parentValue;
            this.heapifyUp(p);
        }
    }

    private getParentIdx(idx: number): number {
        return Math.floor((idx - 1) / 2);
    }

    private getLeftIdx(idx: number): number {
        return idx * 2 + 1;
    }

    private getRightIdx(idx: number): number {
        return idx * 2 + 2;
    }
}
