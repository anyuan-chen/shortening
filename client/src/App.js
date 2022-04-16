import { useState } from "react";
import { ThemeProvider } from "styled-components";
import Button from "./components/ui/Button";
import Header from "./components/ui/Header";
import TextInput from "./components/ui/TextInput";
import theme from "./theme";
function App() {
  const [s, setS] = useState("");

  return (
    <ThemeProvider theme={theme}>
      <Button onClick={() => console.log("hi")}>add</Button>
      <Header type="h1">dfdsdsf</Header>
      <Header type="h2">dfdsdsf</Header>
      <TextInput
        variant="large"
        value={s}
        onChange={(event) => setS(event.target.value)}
      ></TextInput>
      <TextInput
        variant="small"
        value={s}
        onChange={(event) => setS(event.target.value)}
      ></TextInput>
    </ThemeProvider>
  );
}

export default App;
