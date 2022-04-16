import styled from "styled-components";
import React from "react";
import theme from "../../theme";
const StyledButton = styled.button`
  background: #c0daf4;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0.5rem 2rem;
`;

export default function Button({ children, onClick }) {
  return (
    <StyledButton onClick={onClick} style={theme.b1}>
      {children}
    </StyledButton>
  );
}
