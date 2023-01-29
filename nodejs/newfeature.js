// feature: string panding
console.log('hello'.padEnd(10, '*'));
console.log('hello'.padStart(10, '*'));

// feature: promise finally
// 1. promise.reject不能触发finally, throw expection才可以触发finally
// 2. finally 方法可以有多个
const execSuccess = async () => {
  console.log('async run');
}
const execFailure = async () => {
  return Promise.resolve(new Error('debug error'));
}
const execExcpetion = async () => {
  throw new Error('debug error');
}
// execSuccess()
//   .then(() => {
//     return execFailure();
//   })
//   .catch(err => {
//     console.error('catch the error:', err)
//   })
//   .finally(() => {
//     console.log('run the promise finally function 1')
//   })
//   .finally(() => {
//     console.log('run the promise finally function 2')
//   });
// execSuccess()
//   .then(() => {
//     return execExcpetion();
//   })
//   .catch(err => {
//     console.error('catch the error:', err)
//   })
//   .finally(() => {
//     console.log('run the promise finally function')
//   })

// feature: 正则表达式反向
// (?=pattern) p 前面(位置), 零宽正向肯定断言(zero-width positive lookahead assertion)
// (?!pattern) 除了 p 前面(位置), 零宽正向否定断言(zero-width negative lookahead assertion)
// (?<=pattern) p 后面(位置), 零宽反向肯定断言(zero-width positive lookbehind assertion)
// (?<!pattern) 除了 p 后面(位置), 零宽反向否定断言(zero-width negative lookbehind assertion)
let regexp = /just(?=java)/; // java前面位置匹配just
console.log('regexp: ', regexp.exec('1justjavac')[0])
console.log('regexp: ', regexp.test('jjustjavac'))
regexp = /php(?!java)/; // java后面位置匹配php
console.log('regexp: ', regexp.exec('javacphp1234')[0])
console.log('regexp: ', regexp.test('java456php123'))
console.log('regexp: ', regexp.test('java456phhp123'))
console.log('regexp: ', regexp.test('phpjava456'))

// feature: Array.flat, Array.flatMap
console.log([[1, 2], [3, 4]].flat(Infinity));
console.log([[1, 2], [3, 4]].flatMap((element, index, array) => {
  return [element[0], element[1], element[0] + element[1]]
}));

// feature:  transforms a list of key-value pairs into an objec
const entries = new Map([
  ['foo', 'bar'],
  ['baz', 42]
]);

const obj = Object.fromEntries(entries);

console.log(obj);
console.log(Object.entries(obj));
// Expected output: Object { foo: "bar", baz: 42 }

// feature: 空值处理
// 只有undefined和null可判定
console.log(undefined ?? "空值处理默认值: 1")
console.log(null ?? "空值处理默认值: 2")
console.log(false ?? "空值处理默认值: 3")
console.log(0 ?? "空值处理默认值: 4")

// feature: Promise.allSettled
// 方法以 promise 组成的可迭代对象作为输入，并且返回一个 Promise 实例。
// 当输入的所有 promise 都已敲定时（包括传递空的可迭代类型），返回的 promise 将兑现，并带有描述每个 promsie 结果的对象数组。
// Promise.allSettled() 方法是 promise 并发性方法的其中之一。在你有多个不依赖于彼此成功完成的异步任务时，或者你总是想知道每个 promise 的结果时，使用 Promise.allSettled() 。

// 相比之下，如果任务相互依赖，或者如果你想立即拒绝其中任何任务，Promise.all() 返回的 Promise 可能更合适
const promise1 = Promise.resolve(3);
const promise2 = new Promise((resolve, reject) => setTimeout(reject, 100, 'foo'));
const promises = [promise1, promise2];

Promise.allSettled(promises).
  then((results) => results.forEach((result) => {
    console.log(result)
  }));

// Expected output:
// "fulfilled"
// "rejected"

// feature: 数字分隔
console.log('数字分隔不影响值: ', 1_000_000_000 === 1000000000); // true

// feature: 逻辑赋值语法
// x &= y, 等价于 x && (x = y)
// x ||= y, 等价于 x || (x = y)
const a = { duration: 50, title: '' };

a.duration ||= 10;
console.log('逻辑赋值语法1: ', a.duration);
// Expected output: 50

a.title ||= 'title is empty.';
console.log('逻辑赋值语法2:', a.title);
// Expected output: "title is empty"

// feature: Promise
// 1. Promise.all, promise 数组中每个promise都resolve 一旦有任何一个promise reject, 则全部reject
// 2. Promise.allSettled, promise数组中所有的promise都resolve或者reject, 所有的结果组成数组异步返回
// 3. Promise.any, promise数组中任何一个promise resolve就结束, 目的是获取首个resolve的promise
// 4. Promise.race, promise数组中任何一个promise resolve/reject则结束