// 闭包（closure）是一个函数以及其捆绑的周边环境状态（lexical environment，词法环境）的引用的组合。换而言之，闭包让开发者可以从内部函数访问外部函数的作用域。在 JavaScript 中，闭包会随着函数的创建而被同时创建。

function generateAdd() {
  let i = 0;
  function add(n) {
    i += n;
    return i;
  }

  return add;
}

const c1 = generateAdd();
const c2 = generateAdd();
console.log(c1(1));
console.log(c2(3));