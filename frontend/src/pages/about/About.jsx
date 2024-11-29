import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  Box,
  Container,
  Typography,
} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";

import React from "react";

export default function About() {
  return (
    <Box>
      <Container>
        <Box sx={{ display: "flex", justifyContent: "center" }}>
          <Typography component="h1" variant="h2">
            О нас
          </Typography>
        </Box>
        <Box sx={{ width: "100%" }}>
          <img style={{ width: "100%" }} src="/Line 1.png" alt="line" />
        </Box>
        <Box
          sx={{
            mt: 5,
            mb: 5,
            display: "flex",
            flexDirection: "column",
          }}
        >
          <Box sx={{ width: "100%" }}>
            <img style={{ width: "100%" }} src="/about.png" alt="" />
          </Box>
          <Box
            sx={{ display: "flex", flexDirection: "column", gridGap: "40px" }}
          >
            <Typography
              sx={{
                textAlign: "center",
                fontSize: { xs: "20px", md: "30px" },
                fontWeight: "medium",
              }}
              component="h2"
            >
              Средства реабилитации, Товары медицинского назначения и
              медицинская техника
            </Typography>
            <Accordion
              sx={{
                background: "#90E0D4",
                color: "#fff",
              }}
            >
              <AccordionSummary
                expandIcon={
                  <ExpandMoreIcon
                    fontSize="medium"
                    sx={{
                      color: "#fff",
                    }}
                  />
                }
                sx={{ fontSize: "20px" }}
              >
                Здоровье
              </AccordionSummary>
              <AccordionDetails
                sx={{
                  maxHeight: 200,
                  overflow: "auto",
                }}
              >
                <Typography variant="h5" component="h2">
                  Хрупкая вещь, его нужно поддерживать и восстанавливать. Людям
                  с хроническими заболеваниями, в периоды послеоперационной
                  реабилитации, при уходе за больными на дому требуются
                  специализированные изделия медицинского назначения: но где их
                  купить, если в стандартный ассортимент аптек эти позиции не
                  входят?
                </Typography>
              </AccordionDetails>
            </Accordion>
            <Accordion
              sx={{
                background: "#90E0D4",
                color: "#fff",
              }}
            >
              <AccordionSummary
                expandIcon={
                  <ExpandMoreIcon
                    fontSize="medium"
                    sx={{
                      color: "#fff",
                    }}
                  />
                }
                sx={{ fontSize: "20px" }}
              >
                Наш опыт работы с 2000 года. 
              </AccordionSummary>
              <AccordionDetails
                sx={{
                  maxHeight: 200,
                  overflow: "auto",
                }}
              >
                <Typography variant="h5" component="h2">
                  Мы предлагаем большой выбор СРЕДСТВ РЕАБИЛИТАЦИИ ( коляски
                  инвалидные, калоприемники, катетеры, уроприемники и другие
                  средства по уходу)
                </Typography>
              </AccordionDetails>
            </Accordion>
            <Accordion
              sx={{
                background: "#90E0D4",
                color: "#fff",
              }}
            >
              <AccordionSummary
                expandIcon={
                  <ExpandMoreIcon
                    fontSize="medium"
                    sx={{
                      color: "#fff",
                    }}
                  />
                }
                sx={{ fontSize: "20px" }}
              >
                Наши преимущества
              </AccordionSummary>
              <AccordionDetails
                sx={{
                  maxHeight: 200,
                  overflow: "auto",
                }}
              >
                <Typography variant="h5" component="h2">
                  Предоставим консультации менеджеров с медицинским образованием
                  <br></br>
                  Доставим ваш заказ или отгрузим его со склада магазина
                  самостоятельно
                </Typography>
              </AccordionDetails>
            </Accordion>
          </Box>
        </Box>
      </Container>
    </Box>
  );
}
