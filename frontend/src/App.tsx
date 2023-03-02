import type { Component } from "solid-js";

import logo from "./logo.svg";
import styles from "./App.module.css";
import { getPoem } from "./api/getPoem";
import Card from "@suid/material/Card";
import { CssBaseline, ThemeProvider } from "@suid/material";
import { theme } from "./theme";
import { MemeCard } from "./components/MemeCard/MemeCard";
import { Memes } from "./containers/Memes";


const App: Component = () => {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <main class={styles["main"]}>
        <Memes />
      </main>
    </ThemeProvider>
  );
};

export default App;
