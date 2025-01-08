import { AppBar, Toolbar, Typography, Button } from "@mui/material";

const NavBar = () => {
  return (
    <AppBar position="static">
      <Toolbar>
        <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
          Административный дешборд
        </Typography>
        <Button color="inherit">Категории</Button>
        <Button color="inherit">Продукты</Button>
        <Button color="inherit">Пользователи</Button>
      </Toolbar>
    </AppBar>
  );
};
export default NavBar;
