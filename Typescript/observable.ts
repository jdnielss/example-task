type ObservableChange<T> = (oldValue: T, newValue: T) => void;

interface User {
  name: string;
  age: number;
}

const user = {
  name: "Andy",
  age: 20,
};
