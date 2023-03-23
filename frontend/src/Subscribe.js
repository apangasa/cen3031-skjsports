import React from "react";
import ReactDOM from "react-dom/client";
import { useState } from "react";

function SubscribeForm() {
  const [input, setInput] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    async function returnEmail() {
      const subscribe = document.getElementById("Subscribe!").value;
      const response = await fetch(
        "http://localhost:8080/subscribe?subscribe=" + input
      );

      const json = await response.json();
    }
  };

  return (
    <form>
      <label>
        Subscribe!
        <input type="text" email="Email:" />
      </label>
      <input type="submit" value="Submit" />
    </form>
  );
}
const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(<SubscribeForm />);
