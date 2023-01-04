export default function bs_list(haystack: number[], needle: number): boolean {
    let high = haystack.length;
    let low = 0;
    const binary_search = (
        arr: number[],
        hi: number,
        lo: number,
        val: number,
    ): boolean => {
        if (lo > hi) {
            return false 
        }
        const mid = Math.floor(lo + (hi - lo) / 2);
        if (arr[mid] === val) {
            return true;
        } else if (val > arr[mid]) {
            return binary_search(arr, hi, mid + 1, val);
        }
        return binary_search(arr, mid - 1, lo, val);
    };

    // loop version
    // while (low < high) {
    //     const mid = Math.floor(low + (high - low) / 2);
    //
    //     if (haystack[mid] === needle) {
    //         return true;
    //     } else if (haystack[mid] < needle) {
    //         low = mid + 1;
    //     } else {
    //         high = mid - 1;
    //     }
    // }

    return binary_search(haystack, high, low, needle);
}
