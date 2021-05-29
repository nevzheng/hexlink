import Card from "react-bootstrap/Card";
import CardGroup from "react-bootstrap/CardGroup";

const Information: React.FC = () => {
  return (
    <div>
      <h3>Implementation</h3>
      <CardGroup className="p-3">
        <Card style={{ width: "18rem" }}>
          <Card.Body>
            <Card.Title>Backend</Card.Title>
            <Card.Text>
              The backend is implemented in Golang using a Hexagonal
              Archtecture. This architectural pattern results in clean code that
              is easier to modify and extend.
            </Card.Text>
          </Card.Body>
        </Card>
        <Card style={{ width: "18rem" }}>
          <Card.Body>
            <Card.Title>Frontend</Card.Title>
            <Card.Text>
              The Frontend client that you are interacting with at the moment is
              implemented in React with Contexts and Typescript, and Boostrap.
            </Card.Text>
          </Card.Body>
        </Card>
      </CardGroup>
    </div>
  );
};

export default Information;
