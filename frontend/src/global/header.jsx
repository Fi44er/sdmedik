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
  Accordion,
  AccordionSummary,
  AccordionDetails,
  Paper,
} from "@mui/material";
import { withStyles } from "@mui/styles";
import React, { useState } from "react";
import SearchIcon from "@mui/icons-material/Search";
import MenuIcon from "@mui/icons-material/Menu";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";

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
  padding: "20px 40px",
  boxShadow: "0px 4px 10px rgba(0, 0, 0, 0.1)",
  [theme.breakpoints.down("lg")]: {
    // Используйте медиа-запрос для скрытия на мобильных устройствах
    display: "none",
  },
}));
const Search = styled("Box")(({ theme }) => ({
  height: "53px",
  width: { xs: "100%", sm: "100%", md: "50%" },
  border: "3px solid #87EBEB",
  borderRadius: "10px",
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
  const [menuLk, setMenuLk] = React.useState(null);
  const openLk = Boolean(menuLk);

  const handleRegionChange = (event) => {
    setSelectedRegion(event.target.value);
  };

  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };
  const handleClickLk = (event) => {
    setMenuLk(event.currentTarget);
  };
  const handleCloseLk = () => {
    setMenuLk(null);
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

  const navItems = [
    { text: "Доставка", href: "/delivery" },
    { text: "Реквизиты", href: "/deteils" },
    {
      text: "Возврат",
      href: "/returnpolicy",
    },
    {
      text: "О нас",
      href: "/about",
    },
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
                  src="/public/Logo_Header.png"
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
                <img src="/public/Phone.png" alt="" />
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
                  height: "54px",
                  background: `linear-gradient(95.61deg, #A5DED1 4.71%, #00B3A4 97.25%)`,
                  borderRadius: "20px",
                  fontSize: "18px",
                }}
              >
                Каталог
              </Button>
            </Box>
            <Nav>
              {navItems.map((item) => {
                return (
                  <Link
                    underline="hover"
                    sx={{ ml: 2, mr: 2 }}
                    color="black"
                    href={item.href}
                  >
                    {item.text}
                  </Link>
                );
              })}
            </Nav>
            <Box>
              <IconButton
                id="lk-button"
                aria-controls={openLk ? "lk-menu" : undefined}
                aria-haspopup="true"
                aria-expanded={openLk ? "true" : undefined}
                onClick={handleClickLk}
              >
                <img src="/public/Profile.png" alt="" />
              </IconButton>
              <Menu
                id="lk-menu"
                anchorEl={menuLk}
                open={openLk}
                onClose={handleCloseLk}
                MenuListProps={{
                  "aria-labelledby": "lk-button",
                }}
              >
                <MenuItem onClick={handleCloseLk}>
                  <Link href="/auth">Войти</Link>
                </MenuItem>
                <MenuItem onClick={handleCloseLk}>
                  <Link href="/register">Регистрация</Link>
                </MenuItem>
              </Menu>
              <IconButton
                onClick={(e) => {
                  e.preventDefault();
                  window.location.href = "/basket/:id";
                }}
              >
                <img src="/public/basket_header.png" alt="" />
              </IconButton>
            </Box>
          </Box>
        </StyledToolbar>
      </Container>
      {/* Burger menu */}
      <Toolbar
        sx={{
          display: { xs: "flex", sm: "flex", md: "none" },
          flexDirection: "column",
          gridGap: "20px",
        }}
      >
        <Box
          sx={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
            width: "100%",
          }}
        >
          <Link href="/">
            <img
              style={{ width: "60px" }}
              src="/public/Logo_Header.png"
              alt="Logo"
            />
          </Link>
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
          <IconButton
            edge="start"
            color="inherit"
            aria-label="menu"
            onClick={toggleDrawer(true)}
            sx={{
              display: {
                xs: "block",
                sm: "none",
                color: "#26BDB8",
              },
            }}
          >
            <MenuIcon fontSize="large" />
          </IconButton>
        </Box>
        <Box
          sx={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
            width: "100%",
          }}
        >
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
          <IconButton
            onClick={(e) => {
              e.preventDefault();
              window.location.href = "/basket/:id";
            }}
          >
            <img
              style={{ width: "50px", height: "50px" }}
              src="/public/basket_header.png"
              alt=""
            />
          </IconButton>
        </Box>
      </Toolbar>
      <Drawer anchor="left" open={drawerOpen} onClose={toggleDrawer(false)}>
        <Box
          sx={{ width: 250 }}
          role="presentation"
          onClick={toggleDrawer(false)}
          onKeyDown={toggleDrawer(false)}
        >
          <List>
            {navItems.map((item) => {
              return (
                <ListItem>
                  <Link
                    underline="hover"
                    sx={{ color: "#26BDB8", ml: 2 }}
                    color="black"
                    href={item.href}
                  >
                    {item.text}
                  </Link>
                </ListItem>
              );
            })}
            <Button
              id="basic-button"
              variant="contained"
              onClick={(e) => {
                e.preventDefault();
                window.location.href = "/catalog";
              }}
              sx={{
                width: "70%",
                height: "40px",
                background: `linear-gradient(95.61deg, #A5DED1 4.71%, #00B3A4 97.25%)`,
                borderRadius: "20px",
                fontSize: "15px",
                ml: 2,
                mt: 2,
              }}
            >
              Каталог
            </Button>
          </List>
        </Box>
      </Drawer>
    </AppBar>
  );
}