import React, { Component } from 'react';
import { Navbar, NavbarBrand, NavItem, Nav } from 'reactstrap';

export default class Navigation extends Component {
  render() {
    return (
      <Navbar color="light" light expand="md">
        <NavbarBrand>So sánh hiệu năng MongoDB vs MySQL</NavbarBrand>
        <Nav className="ml-auto" navbar>
          <NavItem>
            Have a nice day!
          </NavItem>
          
        </Nav>
      </Navbar>
    )
  }
}