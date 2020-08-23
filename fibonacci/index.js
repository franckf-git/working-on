const number = process.argv[2]

const addTwo = (num1,num2) => {
    return num1 + num2
}

const fibo = (size) => {
    let suite = [1,1]
    while (suite.length < size){
        let nMinus1 = suite[suite.length - 1]
        let nMinus2 = suite[suite.length - 2]
        let current = addTwo(nMinus1,nMinus2)
        suite.push(current)
    }
    console.log(suite)
    return suite
}

fibo(number)

module.exports = { addTwo, fibo, binet }

