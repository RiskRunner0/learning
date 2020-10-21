import * as React from "react";
import * as ReactDom from "react-dom";

const App: React.FunctionComponent = () => {
    return <h1>My React and Typescript App!</h1>
}

ReactDom.render(<App />, document.getElementById("root") as HTMLElement);