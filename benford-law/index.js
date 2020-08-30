
/**
 * law of anomalous numbers
 * @param {number} c
 * @returns {number}
 */

const law = (c) => {
    return (Math.log10(c + 1) - Math.log10(c)) * 100
}

/**
 * compute a range of number and return an array of %
 * @param {number} start
 * @param {number} end
 * @returns {Array}
 */

const rangeLaw = (start, end) => {
    let range = []
    for (let index = start; index <= end; index++) {
        range.push(index)
    }

    const rangeByLaw = range.map((number) => {
        const rounded = law(number).toString().split(".")
        const roundedInt = rounded[0]
        const roundedDec = rounded[1].slice(0, 2)
        const roundedFull = parseFloat([roundedInt, roundedDec].join("."))
        return roundedFull
    })

    return rangeByLaw
}

console.log(rangeLaw(1, 9))