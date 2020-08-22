
const processNumber = (num) => {
    if (num % 2 === 0) {
        return num / 2
    } else {
        return (num * 3) + 1
    }
}

module.exports = processNumber
