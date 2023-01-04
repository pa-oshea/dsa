export default class ArrayList<T> {
    public length: number;
	private data: Array<T>

    constructor() {
        this.length = 0;
		this.data = new Array(10)
    }

    prepend(item: T): void {
		
	}
    insertAt(item: T, idx: number): void {}
    append(item: T): void {}
    remove(item: T): T | undefined {}
    get(idx: number): T | undefined {}
    removeAt(idx: number): T | undefined {}
}
