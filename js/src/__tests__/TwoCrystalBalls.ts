import two_crystal_balls from "@code/TwoCrystalBalls";

test("two crystal balls", function () {
    let idx = Math.floor(Math.random() * 10000);
    const data = new Array(10000).fill(false);
    const data2 = new Array(10000).fill(false);
    const end = data2.length - 1
    data2[end] = true;
    for (let i = idx; i < 10000; ++i) {
        data[i] = true;
    }

    expect(two_crystal_balls(data)).toEqual(idx);
    expect(two_crystal_balls(data2)).toEqual(end);
    expect(two_crystal_balls(new Array(821).fill(false))).toEqual(-1);
});
