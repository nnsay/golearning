function Person(name, age) {
  this.name = name;
  this.age = age;
}
Person.prototype.sayHello = function () {
  console.log(`I'am ${this.name}, a person.`)
}
Person.prototype.introduce = function () {
  console.log(`I'am ${this.name}, ${this.age} years old.`)
}

function Student(name, age, grade) {
  // Person.apply(this, [name, age]);
  Person.call(this, name, age);
  this.grade = grade
}
// Student.prototype = Person.prototype;
Student.prototype = Object.create(Person.prototype);
Student.prototype.sayHello = function () {
  console.log(`I'am ${this.name}, a student.`)
}

const s = new Student('Jimmy', 13, 9);
console.log({ name: s.name, age: s.age, grade: s.grade })
s.sayHello();
console.log(1, s.__proto__)
console.log(2, s.__proto__.__proto__)
s.introduce();