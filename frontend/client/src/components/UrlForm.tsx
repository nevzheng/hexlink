import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";

const UrlForm: React.FC = () => {
  return (
    <div>
      <Form>
        <Form.Group controlId="formURLInput">
          <Form.Label>URL to Shorten</Form.Label>
          <Form.Control type="text" placeholder="http://shorten.me" />
        </Form.Group>
        <Button className="mb-3" variant="primary" type="submit">
          Submit
        </Button>
      </Form>
    </div>
  );
};

export default UrlForm;
