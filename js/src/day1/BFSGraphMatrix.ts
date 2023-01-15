import Queue from "./Queue";

export default function bfs(
    graph: WeightedAdjacencyMatrix,
    source: number,
    needle: number,
): number[] | null {
    const seen = new Array(graph.length).fill(false);
    const prev = new Array(graph.length).fill(-1);

    seen[source] = true;

    const queue = new Queue<number>();
    queue.enqueue(source);

    do {
        const curr = queue.deque();
        // if (curr === undefined) {
        //     break;
        // }
        if (curr === needle) {
            break;
        }

        const adjs = graph[curr!];

        for (let i = 0; i < adjs.length; i++) {
            if (adjs[i] === 0) {
                continue;
            }
            if (seen[i]) {
                continue;
            }

            seen[i] = true;
            prev[i] = curr;
            queue.enqueue(i);
        }
    } while (queue.length);

    if (prev[needle] === -1) {
        return null;
    }
    // build it backwards
    let curr = needle;
    const out: number[] = [];

    while (prev[curr] !== -1) {
        out.push(curr);
        curr = prev[curr];
    }

    return [source].concat(out.reverse());
}
