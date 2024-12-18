import { Box, Container, List, ListItem, Typography } from "@mui/material";
import { Helmet } from "react-helmet";
import React from "react";

export default function Deteils() {
  return (
    <Box sx={{ mb: 5 }}>
      <Helmet>
        <title>Реквизиты компании СД-МЕД - Полная информация о компании</title>
        <meta
          name="description"
          content="Узнайте полные реквизиты компании СД-МЕД, включая ИНН, КПП, адрес и контактные данные. Мы предоставляем качественные медицинские услуги."
        />
        <meta
          name="keywords"
          content="реквизиты СД-МЕД, ИНН СД-МЕД, контактные данные СД-МЕД, адрес СД-МЕД"
        />
      </Helmet>
      <Container>
        <Box sx={{ display: "flex", justifyContent: "center" }}>
          <Typography component="h1" variant="h2" sx={{ textAlign: "center" }}>
            Реквизиты
          </Typography>
        </Box>
        <Box sx={{ width: "100%" }}>
          <img
            style={{ width: "100%" }}
            src="/Line 1.png"
            alt="Линия разделения"
          />
        </Box>
        <Box>
          <Typography sx={{ textAlign: "center" }} component="h2" variant="h4">
            ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ «СД-МЕД»
          </Typography>
          <Box>
            <List>
              <ListItem>
                <Typography variant="h6">
                  ИНН 5609198444, КПП 560901001 ОГРН 1225600000361 460005,
                  Оренбургская область, г. Оренбург , ул. Шевченко д. 20В, этаж
                  1, офис 1
                </Typography>
              </ListItem>
              <ListItem>
                <Typography variant="h6">БИК 042202824,</Typography>
              </ListItem>
              <ListItem>
                <Typography variant="h6">К/с 30101810200000000824</Typography>
              </ListItem>
              <ListItem>
                <Typography variant="h6">Р/с 40702810529250005622</Typography>
              </ListItem>
              <ListItem>
                <Typography variant="h6">
                  E-mail: Sd2-info@yandex.ru www.sdmedik.ru
                </Typography>
              </ListItem>
            </List>
          </Box>
        </Box>
      </Container>
    </Box>
  );
}
