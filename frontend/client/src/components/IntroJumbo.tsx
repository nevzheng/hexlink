import Jumbotron from "react-bootstrap/Jumbotron";
import Nav from "react-bootstrap/Nav";

const IntroJumbo: React.FC = () => {
  return (
    <div>
      <Jumbotron>
        <h1 className="header">Hexlink</h1>
        <h4>
          <i>A Url Shortening Service</i>
        </h4>
        <Nav>
          <Nav.Link href="https://github.com/nevzheng/hexlink">Github</Nav.Link>
          <Nav.Link href="https://github.com/nevzheng/hexlink">
            Docs (TODO)
          </Nav.Link>
        </Nav>
      </Jumbotron>
    </div>
  );
};

export default IntroJumbo;
