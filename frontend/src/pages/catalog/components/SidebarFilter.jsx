import React, { useState } from "react";
import {
  Box,
  Button,
  Drawer,
  IconButton,
  Slider,
  Typography,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  useMediaQuery,
  useTheme,
  Accordion,
  AccordionSummary,
  AccordionDetails,
  FormControlLabel,
  Radio,
  RadioGroup,
  TextField,
  styled,
  Paper,
} from "@mui/material";
import FilterListIcon from "@mui/icons-material/FilterList";
import CloseIcon from "@mui/icons-material/Close";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";

const CustomTextField = styled(TextField)({
  "& .MuiOutlinedInput-root": {
    "& fieldset": {
      borderColor: "#26BDB8",
    },
    "&:hover fieldset": {
      borderColor: "#26BDB8",
    },
    "&.Mui-focused fieldset": {
      borderColor: "#26BDB8",
    },
  },
});

const SidebarFilter = () => {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down("sm")); // Определяем, мобильное ли устройство
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [priceRange, setPriceRange] = useState([20000, 30000]);
  const [selectedOffer, setSelectedOffer] = useState("");
  const [selectedColor, setSelectedColor] = useState("");
  const [selectedBatteryType, setSelectedBatteryType] = useState("");
  const [selectedMotorPower, setSelectedMotorPower] = useState("");

  const controlProps = (item) => ({
    // checked: selectedValue === item,
    // onChange: handleChange,
    value: item,
    name: "color-radio-button-demo",
    inputProps: { "aria-label": item },
  });

  const toggleDrawer = () => {
    setDrawerOpen(!drawerOpen);
  };

  const handlePriceChange = (event, newValue) => {
    setPriceRange(newValue);
  };

  const handleApplyFilters = () => {
    // Здесь вы можете добавить логику применения фильтров
    toggleDrawer(); // Закрытие меню после применения фильтров
  };

  return (
    <Box sx={{ display: "flex" }}>
      {isMobile && (
        <Box sx={{ mt: 5 }}>
          <Button
            sx={{
              background: "#00B3A4",
              color: "white",
              height: "50px",
              width: "150px",
            }}
            onClick={toggleDrawer}
          >
            Фильрация
            <FilterListIcon />
          </Button>
        </Box>
      )}

      {/* Drawer для мобильной версии */}
      <Drawer
        anchor="left"
        open={drawerOpen}
        onClose={toggleDrawer}
        sx={{ "& .MuiDrawer-paper": { width: "100vw", height: "100vh" } }} // Занимает весь экран
      >
        <Box sx={{ padding: 2, height: "100vh", position: "relative" }}>
          <IconButton
            onClick={toggleDrawer}
            sx={{ position: "absolute", right: 16, top: 16 }}
          >
            <CloseIcon />
          </IconButton>
          <Typography variant="h6">Фильтрация</Typography>
          <Box sx={{ mt: 2 }}>
            <Box sx={{ mb: 2 }}>
              <Typography variant="body">Цена</Typography>
            </Box>
            <Box
              sx={{
                display: "flex",
                justifyContent: "center",
                gridGap: "20px",
              }}
            >
              <Box>
                <label>От</label>
                <CustomTextField
                  variant="outlined"
                  sx={{ width: "100%", mt: 2 }}
                />
              </Box>
              <Box>
                <label>До</label>
                <CustomTextField
                  variant="outlined"
                  sx={{ width: "100%", mt: 2 }}
                />
              </Box>
            </Box>
          </Box>
          <FormControl fullWidth sx={{ mt: 2 }}>
            <Accordion>
              <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                Наши предложения
              </AccordionSummary>
              <AccordionDetails
                sx={{
                  maxHeight: 200,
                  overflow: "auto",
                }}
              >
                <RadioGroup
                  aria-labelledby="demo-radio-buttons-group-label"
                  defaultValue="female"
                  name="radio-buttons-group"
                >
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                </RadioGroup>
              </AccordionDetails>
            </Accordion>
          </FormControl>
          <FormControl fullWidth sx={{ mt: 2 }}>
            <Accordion>
              <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                Цвет рамы
              </AccordionSummary>
              <AccordionDetails
                sx={{
                  maxHeight: 200,
                  overflow: "auto",
                }}
              >
                <RadioGroup
                  aria-labelledby="demo-radio-buttons-group-label"
                  defaultValue="female"
                  name="radio-buttons-group"
                >
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                </RadioGroup>
              </AccordionDetails>
            </Accordion>
          </FormControl>
          <FormControl fullWidth sx={{ mt: 2 }}>
            <Accordion>
              <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                Тип аккумулятора
              </AccordionSummary>
              <AccordionDetails
                sx={{
                  maxHeight: 200,
                  overflow: "auto",
                }}
              >
                <RadioGroup
                  aria-labelledby="demo-radio-buttons-group-label"
                  defaultValue="female"
                  name="radio-buttons-group"
                >
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                </RadioGroup>
              </AccordionDetails>
            </Accordion>
          </FormControl>
          <FormControl fullWidth sx={{ mt: 2 }}>
            <Accordion>
              <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                Мощность электромотора
              </AccordionSummary>
              <AccordionDetails
                sx={{
                  maxHeight: 200,
                  overflow: "auto",
                }}
              >
                <RadioGroup
                  aria-labelledby="demo-radio-buttons-group-label"
                  defaultValue="female"
                  name="radio-buttons-group"
                >
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                  <FormControlLabel
                    value="female"
                    control={
                      <Radio
                        {...controlProps("e")}
                        sx={{
                          color: "#00B3A4",
                          "&.Mui-checked": {
                            color: "#00B3A4",
                          },
                        }}
                      />
                    }
                    label="Female"
                  />
                </RadioGroup>
              </AccordionDetails>
            </Accordion>
          </FormControl>
          <Button
            variant="contained"
            sx={{
              mt: 3,
              background: `linear-gradient(91.54deg, #9FE1D2 -43.68%, #2AC8BB 142.9%, #14B8A9 142.9%)`,
              borderRadius: "50px",
              height: "50px",
              width: "100%",
            }}
            // onClick={toggleDrawer}
          >
            Применить фильтры
          </Button>
          <Button
            variant="contained"
            sx={{
              mt: 3,
              background: `#BEF4F0`,
              borderRadius: "50px",
              height: "50px",
              color: "black",
              width: "100%",
            }}
          >
            Очистить
          </Button>
        </Box>
      </Drawer>

      {/* Меню фильтров для десктопа всегда отображается */}
      <Box
        sx={{
          width: "250px",
          display: isMobile ? "none" : "block", // Скрыть на мобильной версии
          padding: 2,
        }}
      >
        <Box
          sx={{
            width: "100%",
            background: "#00B3A4",
            height: "40px",
            borderRadius: "50px",
          }}
        >
          <Typography
            variant="h6"
            sx={{
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
              color: "white",
            }}
          >
            Фильтрация
          </Typography>
        </Box>
        <Box sx={{ mt: 2 }}>
          <Box sx={{ mb: 2 }}>
            <Typography variant="body">Цена</Typography>
          </Box>
          <Box
            sx={{
              display: "flex",
              justifyContent: "center",
              gridGap: "20px",
            }}
          >
            <Box>
              <label>От</label>
              <CustomTextField
                variant="outlined"
                sx={{ width: "100%", mt: 2 }}
              />
            </Box>
            <Box>
              <label>До</label>
              <CustomTextField
                variant="outlined"
                sx={{ width: "100%", mt: 2 }}
              />
            </Box>
          </Box>
        </Box>
        <FormControl fullWidth sx={{ mt: 2 }}>
          <Accordion>
            <AccordionSummary expandIcon={<ExpandMoreIcon />}>
              Наши предложения
            </AccordionSummary>
            <AccordionDetails
              sx={{
                maxHeight: 200,
                overflow: "auto",
              }}
            >
              <RadioGroup
                aria-labelledby="demo-radio-buttons-group-label"
                defaultValue="female"
                name="radio-buttons-group"
              >
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
              </RadioGroup>
            </AccordionDetails>
          </Accordion>
        </FormControl>
        <FormControl fullWidth sx={{ mt: 2 }}>
          <Accordion>
            <AccordionSummary expandIcon={<ExpandMoreIcon />}>
              Цвет рамы
            </AccordionSummary>
            <AccordionDetails
              sx={{
                maxHeight: 200,
                overflow: "auto",
              }}
            >
              <RadioGroup
                aria-labelledby="demo-radio-buttons-group-label"
                defaultValue="female"
                name="radio-buttons-group"
              >
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
              </RadioGroup>
            </AccordionDetails>
          </Accordion>
        </FormControl>
        <FormControl fullWidth sx={{ mt: 2 }}>
          <Accordion>
            <AccordionSummary expandIcon={<ExpandMoreIcon />}>
              Тип аккумулятора
            </AccordionSummary>
            <AccordionDetails
              sx={{
                maxHeight: 200,
                overflow: "auto",
              }}
            >
              <RadioGroup
                aria-labelledby="demo-radio-buttons-group-label"
                defaultValue="female"
                name="radio-buttons-group"
              >
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
              </RadioGroup>
            </AccordionDetails>
          </Accordion>
        </FormControl>
        <FormControl fullWidth sx={{ mt: 2 }}>
          <Accordion>
            <AccordionSummary expandIcon={<ExpandMoreIcon />}>
              Мощность электромотора
            </AccordionSummary>
            <AccordionDetails
              sx={{
                maxHeight: 200,
                overflow: "auto",
              }}
            >
              <RadioGroup
                aria-labelledby="demo-radio-buttons-group-label"
                defaultValue="female"
                name="radio-buttons-group"
              >
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
                <FormControlLabel
                  value="female"
                  control={
                    <Radio
                      {...controlProps("e")}
                      sx={{
                        color: "#00B3A4",
                        "&.Mui-checked": {
                          color: "#00B3A4",
                        },
                      }}
                    />
                  }
                  label="Female"
                />
              </RadioGroup>
            </AccordionDetails>
          </Accordion>
        </FormControl>
        <Button
          variant="contained"
          sx={{
            mt: 3,
            background: `linear-gradient(91.54deg, #9FE1D2 -43.68%, #2AC8BB 142.9%, #14B8A9 142.9%)`,
            borderRadius: "50px",
            height: "50px",
            width: "100%",
          }}
          // onClick={toggleDrawer}
        >
          Применить фильтры
        </Button>
        <Button
          variant="contained"
          sx={{
            mt: 3,
            background: `#BEF4F0`,
            borderRadius: "50px",
            height: "50px",
            color: "black",
            width: "100%",
          }}
        >
          Очистить
        </Button>
      </Box>
    </Box>
  );
};

export default SidebarFilter;
