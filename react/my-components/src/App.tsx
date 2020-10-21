import React from "react";
import logo from "./logo.svg";
import "./App.css";
import Confirm from "./Confirm";

interface IState {
  confirmOpen: boolean;
  confirmMessage: string;
  confirmVisible: boolean;
  countDown: number;
}

class App extends React.Component<{}, IState> {
  private timer: number = 0;

  constructor(props: {}) {
    super(props);
    this.state = {
      confirmOpen: false,
      confirmMessage: "Please hit the confirm button",
      confirmVisible: true,
      countDown: 10,
    };
  }

  public render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.tsx</code> and save to reload.
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </header>
        <p>{this.state.confirmMessage}</p>
        {this.state.confirmVisible && (
          <button onClick={this.handleConfirmClick}>Confirm</button>
        )}
        <Confirm
          open={this.state.confirmOpen}
          title="React and Typescript"
          content="Are you sure you want to learn React and Typescript?"
          cancelCaption="No way"
          okCaption="Yes please!"
          onCancelClick={this.handleCancelConfirmClick}
          onOkClick={this.handleOkConfirmClick}
        />
      </div>
    );
  }

  public componentDidMount() {
    this.timer = window.setInterval(() => this.handleTimerTick(), 1000);
  }

  public componentWillUnmount() {
    clearInterval(this.timer);
  }

  public shouldComponentUpdate(nextProps: {}, nextState: IState): boolean
  {
    console.log("shouldComponentUpdate", nextProps, nextState);
    return false;
  }

  public static getDerivedStateFromProps(props: {}, state: IState) {
    console.log("getDerivedStateFromProps", props, state);
    return null;
  }

  private handleTimerTick() {
    this.setState(
      {
        confirmMessage: `Please hit the confirm button ${this.state.countDown} secs to go`,
        countDown: this.state.countDown - 1,
      },
      () => {
        if (this.state.countDown <= 0) {
          clearInterval(this.timer);
          this.setState({
            confirmMessage: "Too late to confirm!",
            confirmVisible: false,
          });
        }
      }
    );
  }

  private handleCancelConfirmClick = () => {
    this.setState({
      confirmOpen: false,
      confirmMessage: "Take a break, I'm sure you will later ...",
    });
    clearInterval(this.timer);
  };

  private handleOkConfirmClick = () => {
    this.setState({
      confirmOpen: false,
      confirmMessage: "Cool, carry on reading!",
    });
    clearInterval(this.timer);
  };

  private handleConfirmClick = () => {
    this.setState({ confirmOpen: true });
    clearInterval(this.timer);
  };
}

export default App;
