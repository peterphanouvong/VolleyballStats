const counterButton = document.querySelector("#counter-button");
const counterValue = document.querySelector("#counter-value");

counterButton.addEventListener("click", async () => {
  console.log("clicked button");
  const data = await incrementCounter();

  counterValue.textContent = data.Data;
});

const incrementCounter = async () => {
  const res = await fetch("http://localhost:10000/counter/increment", {
    headers: {
      "Content-Type": "application/json",
    },
    mode: "no-cors",
  });
  const data = await res.json();
  return data;
};
