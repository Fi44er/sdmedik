import {
  AppBar,
  Avatar,
  Box,
  Button,
  InputBase,
  Link,
  MenuItem,
  styled,
  Toolbar,
  Typography,
  IconButton,
  Drawer,
  List,
  ListItem,
  ListItemText,
  Select,
  InputLabel,
  FormControl,
  Menu,
  TextField,
  InputAdornment,
  Autocomplete,
  Input,
  Container,
} from "@mui/material";
import { withStyles } from "@mui/styles";
import React, { useState } from "react";
import SearchIcon from "@mui/icons-material/Search";

const StyledToolbar = styled(Toolbar)(({ theme }) => ({
  display: "flex",
  gridGap: "25px",
  flexDirection: "column",
  position: "relative",
  [theme.breakpoints.down("sm")]: {
    // Используйте медиа-запрос для скрытия на мобильных устройствах
    display: "none",
  },
}));

const Nav = styled("div")(({ theme }) => ({
  width: "max-content",
  background: "#FAFAFA",
  borderRadius: "15px",
  display: "flex",
  alignItems: "center",
  padding: "20px 60px",
  boxShadow: "0px 4px 10px rgba(0, 0, 0, 0.1)",
  [theme.breakpoints.down("lg")]: {
    // Используйте медиа-запрос для скрытия на мобильных устройствах
    display: "none",
  },
}));
const Search = styled("Box")(({ theme }) => ({
  height: "53px",
  width: "40%",
  border: "3px solid #87CEEB",
  borderRadius: "30px",
  color: "black",
  display: "flex",
  alignItems: "center",
  paddingLeft: "25px",
}));
const regions = [
  "г. Оренбург",
  "г. Москва",
  "г. Санкт-Петербург",
  // Добавьте сюда остальные регионы
  // ...
];

export default function Header() {
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [anchorEl, setAnchorEl] = React.useState(null);
  const open = Boolean(anchorEl);

  const [selectedRegion, setSelectedRegion] = useState(regions[0]);

  const handleRegionChange = (event) => {
    setSelectedRegion(event.target.value);
  };

  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };

  const toggleDrawer = (open) => (event) => {
    if (
      event.type === "keydown" &&
      (event.key === "Tab" || event.key === "Shift")
    ) {
      return;
    }
    setDrawerOpen(open);
  };
  const menuItems = [
    { text: "Доставка", href: "/" },
    { text: "Оплата", href: "/news" },
    { text: "О нас", href: "/help" },
    { text: "Контакты", href: "/contacts" },
  ];
  return (
    <AppBar position="sticky" sx={{ background: "white", p: 1 }}>
      <Container>
        <StyledToolbar>
          <Box
            sx={{
              width: "100%",
              display: "flex",
              justifyContent: "space-around",
              alignItems: "center",
            }}
          >
            <Box sx={{ display: "flex", gridGap: 20 }}>
              <Box
                onClick={(e) => {
                  e.preventDefault();
                  window.location.href = "/";
                }}
                sx={{
                  width: {
                    xs: "150px",
                    sm: "250px",
                    md: "max-content",
                    cursor: "pointer",
                  },
                }}
              >
                <img
                  style={{ width: "100%" }}
                  src="/public/medi_logo_mini.png"
                  alt="logo"
                />
              </Box>
              <Box sx={{ display: "flex", alignItems: "center" }}>
                <Select
                  value={selectedRegion}
                  onChange={handleRegionChange}
                  displayEmpty
                  variant="standard"
                  sx={{ marginLeft: 2 }}
                >
                  {regions.map((region) => (
                    <MenuItem key={region} value={region}>
                      {region}
                    </MenuItem>
                  ))}
                </Select>
              </Box>
            </Box>
            <Search>
              <SearchIcon fontSize="medium" />
              <input
                type="text"
                placeholder="Поиск по товарам"
                style={{
                  width: "100%",
                  height: "90%",
                  border: "none",
                  borderRadius: "30px",
                  fontSize: "18px",
                  outline: "none",
                  marginLeft: "10px",
                }}
              />
            </Search>
            <Box sx={{ display: "flex", alignItems: "center", gridGap: 20 }}>
              <IconButton
                id="basic-button"
                aria-controls={open ? "basic-menu" : undefined}
                aria-haspopup="true"
                aria-expanded={open ? "true" : undefined}
                onClick={handleClick}
              >
                <img src="/public/mobile_phone.png" alt="" />
              </IconButton>
              <Menu
                id="basic-menu"
                anchorEl={anchorEl}
                open={open}
                onClose={handleClose}
                MenuListProps={{
                  "aria-labelledby": "basic-button",
                }}
              >
                <MenuItem onClick={handleClose}>+7 (903) 086 3091</MenuItem>
                <MenuItem onClick={handleClose}>+7 (909) 617 8196</MenuItem>
                <MenuItem onClick={handleClose}>+7 (353) 293 5241</MenuItem>
              </Menu>
              <Typography color="black">olimp1-info@yandex.ru</Typography>
            </Box>
          </Box>
          <Box
            sx={{
              width: "100%",
              display: "flex",
              justifyContent: "space-around",
              alignItems: "center",
            }}
          >
            <Box
              sx={{
                display: "flex,",
                justifyContent: "center",
                alignItems: "center",
              }}
            >
              <Button
                id="basic-button"
                variant="contained"
                onClick={(e) => {
                  e.preventDefault();
                  window.location.href = "/catalog";
                }}
                sx={{
                  width: "200px",
                  height: "64px",
                  background: "#1E90FF",
                  borderRadius: "20px",
                  fontSize: "18px",
                }}
              >
                Каталог
              </Button>
            </Box>
            <Nav>
              {menuItems.map((item) => {
                {
                  return (
                    <Link
                      underline="hover"
                      sx={{ ml: 4, mr: 4 }}
                      color="black"
                      href={item.href}
                    >
                      {item.text}
                    </Link>
                  );
                }
              })}
            </Nav>
            <Box>
              <IconButton>
                <img src="/public/basket.png" alt="" />
              </IconButton>
            </Box>
          </Box>
        </StyledToolbar>
      </Container>
      {/* Burger menu */}
      <Toolbar sx={{ display: { xs: "flex", sm: "flex", md: "none" } }}>
        <Box
          onClick={(e) => {
            e.preventDefault();
            window.location.href = "/";
          }}
          sx={{ flexGrow: 1, cursor: "pointer" }}
        >
          <img style={{ width: "100px" }} alt="Logo" />
        </Box>
        <IconButton
          edge="start"
          color="inherit"
          aria-label="menu"
          onClick={toggleDrawer(true)}
          sx={{
            display: {
              xs: "block",
              sm: "none",
              color: "#C152F0",
            },
          }}
        >
          <img src="/public/MenuIcon.svg" />
        </IconButton>
      </Toolbar>
      <Drawer anchor="left" open={drawerOpen} onClose={toggleDrawer(false)}>
        <Box
          sx={{ width: 250 }}
          role="presentation"
          onClick={toggleDrawer(false)}
          onKeyDown={toggleDrawer(false)}
        >
          <List>
            {menuItems.map((item) => (
              <ListItem button key={item.text} component="a" href={item.href}>
                <ListItemText sx={{ color: "#BE1BF7" }} primary={item.text} />
              </ListItem>
            ))}
          </List>
        </Box>
      </Drawer>
    </AppBar>
  );
}
