type ObservableChange<T> = (oldValue: T, newValue: T) => void;

interface User {
  name: string;
  age: number;
}

const user = {
  name: "Andy",
  age: 20,
};

/**
 * Observable merupakan pattern untuk memanggil fungsi ketika sebuah object berubah nilainya
 * contoh dan penggunaannya: https://codesandbox.io/s/ts-proxy-observable-h1yln?file=/src/index.ts
 * @param obj Object dasar yang menjadi pondasi
 * @param onChange fungsi yang akan dipanggil ketika object dasar berubah nilai property nya
 */
function observable<T>(obj: T, onChange: ObservableChange<T>): T {
  // callback akan dipanggil pertama kali ketika fungsi observable dipanggil
  onChange(obj, obj);

  return new Proxy(obj as Object, {
    set: function (target: T, key, newValue) {
      const oldValue = { ...target };
      // mengubah nilai object berdasarkan property yang diubah nilainya
      Reflect.set(target as Object, key, newValue);
      onChange(oldValue, target);
      return true;
    },
  }) as T;
}

// Contoh penggunaan pada DOM
// const userObservable = observable<User>(user, (oldValue, newValue) => {
//   const heading = document.getElementById("heading");
//   if (heading) {
//     heading.innerHTML = newValue.name;
//   }
// });

// document.getElementById("btn")?.addEventListener("click", () => {
//   userObservable.name = userObservable.name === "Hello" ? "Aww" : "Hello";
// });
