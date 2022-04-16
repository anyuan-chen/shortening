import styled from "styled-components";
import React from "react";
const StyledButton = styled.button`
  background: #c0daf4;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0.5rem 2rem;
  font-style: ${(props) => props.theme.b1.fontStyle};
  font-family: ${(props) => props.theme.b1.fontFamily};
  font-weight: ${(props) => props.theme.b1.fontWeight};
  font-size: ${(props) => props.theme.b1.fontSize};
  line-height: ${(props) => props.theme.b1.lineHeight};
  color: ${(props) => props.theme.b1.color};
`;

export default function Button({ children, onClick }) {
  return <StyledButton onClick={onClick}>{children}</StyledButton>;
}
