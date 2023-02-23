import type { Component } from "solid-js";

import logo from "./logo.svg";
import styles from "./App.module.css";
import { getPoem } from "./api/getPoem";

const App: Component = () => {
  getPoem("test");

  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <img src={logo} class={styles.logo} alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
      </header>
    </div>
  );
};

export default App;
