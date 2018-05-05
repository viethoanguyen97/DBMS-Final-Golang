import React, { Component } from 'react';
import { Form, FormGroup, Label, Button, Input, Col, Row, Fade } from 'reactstrap';

const hostname = "http://192.168.1.93:3001"
const querySetText = ["Chọn một loại xe", "Chọn nhiều loại xe", "Chọn một khách hàng", "Chọn nhiều khách hàng", "Chọn một đơn hàng", "Chọn nhiều đơn hàng", "Lấy tất cả chi tiết của một đơn hàng", "Lấy tất cả chi tiết của một đơn hàng với tên loại xe", "Lấy tất cả chi tiết của một đơn hàng với tên loại xe theo id khách hàng", "Thêm một chi tiết đơn hàng", "Sửa một đơn hàng", "Xoá một chi tiết đơn hàng"]
const methodsSet = ["GET", "GET", "GET", "GET", "GET", "GET", "GET", "GET", "GET", "POST", "PUT", "DELETE"]

export default class Navigation extends Component {
  constructor(props) {
    super(props);
    this.state = {
      queryType: 0,
      dbname: "mongodb",
      idField: 0,
      insertOrderIdField: 0,
      insertCarIdField: 0,
      insertQuantityOrderField: 0,
      carIdField: 0,
      durationTime: 0,
      rows: 0,
      durationTimeFadeIn: false
    }
    this.onChangeQuerryType = this.onChangeQuerryType.bind(this);
    this.onChangeDbms = this.onChangeDbms.bind(this);
    this.onChangeIdField = this.onChangeIdField.bind(this);
    this.fetchData = this.fetchData.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
  }

  onSubmit(e) {
    e.preventDefault();
    let baseUri = hostname + "/api/" + this.state.dbname;
    let uri = baseUri;
    let body = {}
    this.setState({
      durationTime: 0,
      durationTimeFadeIn: false
    })
    switch (this.state.queryType) {
      case 0:
        uri += ("/cars/" + this.state.idField);
        break;
      case 1:
        uri += ("/cars");
        break;
      case 2:
        uri += ("/customers/" + this.state.idField);
        break;
      case 3:
        uri += ("/customers");
        break;
      case 4:
        uri += ("/orders/" + this.state.idField);
        break;
      case 5:
        uri += ("/orders");
        break;
      case 6:
        uri += ("/orders/" + this.state.idField + "/details");
        break;
      case 7:
        uri += ("/orders/" + this.state.idField + "/cars");
        break;
      case 8:
        uri += ("/customers/" + this.state.idField + "/orders");
        break;
      case 9:
        uri += ("/orderdetails");
        body['order_id'] = this.state.insertOrderIdField;
        body['car_id'] = this.state.insertCarIdField;
        body['quantity_order'] = this.state.insertQuantityOrderField;
        break;
      case 10:
        uri += ("/orders/" + this.state.idField);
        break;
      default:
        uri += ("/orders/" + this.state.idField + "?car_id=" + this.state.carIdField);
        break;
    }
    this.fetchData(uri, methodsSet[this.state.queryType], body);
  }

  fetchData(uri, type, body) {
    let options = {}
    options["method"] = methodsSet[this.state.queryType];
    options["headers"] = {
      "Content-Type": "application/json"
    }
    if (methodsSet[this.state.queryType] === "POST") {
      options["body"] = JSON.stringify(body);
    } else if (methodsSet[this.state.queryType] === "PUT") {
      options["body"] = JSON.stringify({
        "customer_id": 3000
      });
    }
    console.log(options);
    fetch(uri, options)
      .then(res => res.json())
      .then((res) => {
        console.log(res);
        if (res.status === 200 || res.status === 400) {
          this.setState({
            durationTime: res.duration,
            durationTimeFadeIn: true,
            rows: res.rows
          })
        }
        return res;
      })
    // console.log(uri, type, body);
  }

  onChangeIdField(e) {
    e.preventDefault();
    if (e.target.value !== "") {
      this.setState({
        [e.target.name]: parseInt(e.target.value, 10)
      })
    } else {
      this.setState({
        [e.target.name]: 0
      })
    }
  }

  onChangeDbms(e) {
    e.preventDefault();
    this.setState({
      dbname: e.target.value
    })
  }

  onChangeQuerryType(e) {
    e.preventDefault();
    // alert(e.target.value);
    this.setState({
      durationTimeFadeIn: false,
      queryType: parseInt(e.target.value, 10)
    })
    // console.log(this.state);
  }

  render() {
    let idFieldLabel = "";
    switch (this.state.queryType) {
      case 0:
        idFieldLabel = "ID Loại xe";
        break;
      case 2: case 8:
        idFieldLabel = "ID Khách hàng";
        break;
      default:
        idFieldLabel = "ID Đơn hàng";
    }
    let idField = (this.state.queryType !== 1 && this.state.queryType !== 3 && this.state.queryType !== 5 && this.state.queryType !== 9)
      ? <FormGroup row>
        <Label sm={3} for="idField">{idFieldLabel}</Label>
        <Col sm={9}>
          <Input type="number" name="idField" id="idField" onChange={this.onChangeIdField} value={this.state.idField} />
        </Col>
      </FormGroup>
      : <div></div>

    let carIdField = (this.state.queryType === 11)
      ? <FormGroup row>
        <Label sm={3} for="carIdField">ID Loại xe</Label>
        <Col sm={9}>
          <Input type="number" name="carIdField" id="carIdField" onChange={this.onChangeIdField} value={this.state.carIdField} />
        </Col>
      </FormGroup>
      : <div></div>

    let insertField = (this.state.queryType === 9)
      ? <div>
        <FormGroup row>
          <Label sm={3} for="insertOrderIdField">ID Đơn hàng</Label>
          <Col sm={9}>
            <Input type="number" name="insertOrderIdField" id="insertOrderIdField" onChange={this.onChangeIdField} value={this.state.insertOrderIdField} />
          </Col>
        </FormGroup>

        <FormGroup row>
          <Label sm={3} for="insertCarIdField">ID Loại xe</Label>
          <Col sm={9}>
            <Input type="number" name="insertCarIdField" id="insertCarIdField" onChange={this.onChangeIdField} value={this.state.insertCarIdField} />
          </Col>
        </FormGroup>

        <FormGroup row>
          <Label sm={3} for="insertQuantityOrderField">Số lượng đặt hàng</Label>
          <Col sm={9}>
            <Input type="number" name="insertQuantityOrderField" id="insertQuantityOrderField" onChange={this.onChangeIdField} value={this.state.insertQuantityOrderField} />
          </Col>
        </FormGroup>
      </div>
      : <div></div>

    return (
      <div>
        <Row>
          <Col sm={{ size: 8, offset: 2 }}>
            <Form>
              <FormGroup row>
                <Label for="queryType" sm={3}>Truy vấn</Label>
                <Col sm={9}>
                  <Input type="select" name="queryType" id="queryType" onChange={this.onChangeQuerryType}>
                    {querySetText.map((queryText, i) => <option key={i} value={i}>{queryText}</option>)}
                  </Input>
                </Col>
              </FormGroup>

              <FormGroup row>
                <Label for="dbms" sm={3}>Hệ quản trị CSDL</Label>
                <Col sm={9}>
                  <Input type="select" name="dbms" id="dbms" onChange={this.onChangeDbms}>
                    <option value="mongodb">MongoDB</option>
                    <option value="mysql">MySQL</option>
                  </Input>
                </Col>
              </FormGroup>

              {idField}
              {carIdField}
              {insertField}

              <Button color="primary" onClick={this.onSubmit}>GO</Button>{' '}
            </Form>
            <br />
            <Fade in={this.state.durationTimeFadeIn}>
              <Button color="success">{this.state.rows} row(s) in {this.state.durationTime} s</Button>
            </Fade>
          </Col>
        </Row>
      </div>
    )
  }
}