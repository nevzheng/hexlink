import Container from "react-bootstrap/Container";

import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";

import Information from "./components/Information";
import IntroJumbo from "./components/IntroJumbo";
import LinkTable from "./components/LinkTable";
import UrlForm from "./components/UrlForm";
import ShortenResult from "./components/ShortenResult";

export const App = () => {
  return (
    <Container className="p-3">
      <IntroJumbo />
      <Row>
        <Col>
          <UrlForm />
        </Col>
        <Col>
          <ShortenResult />
        </Col>
      </Row>
      <hr />
      <LinkTable />
      <hr />
      <Information />
    </Container>
  );
};

export default App;
