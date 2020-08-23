
const processNumber = (num) => {
    if (num % 2 === 0) {
        return num / 2
    } else {
        return (num * 3) + 1
    }
}

const collatzProcess = (entryPoint) => {
    let collatzList = []
    entryPoint = parseInt(entryPoint)
    collatzList.push(entryPoint)

    while (collatzList[collatzList.length - 1] > 1) {
        let lastNumber = collatzList[collatzList.length - 1]
        let numberProcessed = processNumber(lastNumber)
        collatzList.push(numberProcessed)
    }
    return collatzList
}

export { processNumber, collatzProcess }
// module.exports = { processNumber, collatzProcess }
