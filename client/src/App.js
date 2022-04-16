import { ThemeProvider } from "styled-components";
import Button from "./components/ui/Button";
import theme from "./theme";
function App() {
  return (
    <ThemeProvider theme={theme}>
      <Button>add</Button>
    </ThemeProvider>
  );
}

export default App;
