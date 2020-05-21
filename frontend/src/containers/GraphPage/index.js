import React, { useEffect, useState } from "react";
import { Chart } from "react-google-charts";
import {
  useAxiosSafely,
  urlGraphTransactionCategory,
  urlGraphTransactionAccount
} from "../../api";

import { Typography, IconButton } from "@material-ui/core";
import KeyboardArrowRightIcon from "@material-ui/icons/KeyboardArrowRight";
import KeyboardArrowLeftIcon from "@material-ui/icons/KeyboardArrowLeft";

const moment = require("moment");

export default function GraphPage() {
  const [currentMonthViewed, setMonthViewed] = useState(
    moment().startOf("month")
  );
  const setNextMonth = () =>
    setMonthViewed(currentMonthViewed.clone().add(1, "month"));
  const setPreviousMonth = () =>
    setMonthViewed(currentMonthViewed.clone().subtract(1, "month"));
  const getStartOfCurrentMonthAsString = () =>
    currentMonthViewed.format("YYYY-MM-DD");
  const getEndOfCurrentMonthAsString = () =>
    currentMonthViewed
      .clone()
      .endOf("month")
      .format("YYYY-MM-DD");

  const [
    { data: categoryGraphData },
    fetchGraphTransactionCategory
  ] = useAxiosSafely(urlGraphTransactionCategory());
  const [
    { data: accountGraphData },
    fetchGraphTransactionAccount
  ] = useAxiosSafely(urlGraphTransactionAccount());

  useEffect(() => {
    fetchGraphTransactionCategory({
      params: {
        dateTimeStart: getStartOfCurrentMonthAsString(),
        dateTimeEnd: getEndOfCurrentMonthAsString()
      }
    });
    fetchGraphTransactionAccount({
      params: {
        dateTimeStart: getStartOfCurrentMonthAsString(),
        dateTimeEnd: getEndOfCurrentMonthAsString()
      }
    });
  }, [currentMonthViewed]);
  const formattedCategoryGraphData = [["Category", "Amount"]].concat(
    categoryGraphData.map(pieChartSliceData => [
      pieChartSliceData.label,
      pieChartSliceData.quantity
    ])
  );
  const formattedAccountGraphData = [["Account", "Amount"]].concat(
    accountGraphData.map(pieChartSliceData => [
      pieChartSliceData.label,
      pieChartSliceData.quantity
    ])
  );

  const baseGraphOptions = {
    pieHole: 0.5,
    is3D: false
  };
  const categoryGraphOptions = {
    ...baseGraphOptions,
    title: "Transactions by Category"
  };
  const accountGraphOptions = {
    ...baseGraphOptions,
    title: "Transactions by Account"
  };

  return (
    <div>
      <h1>Graph</h1>
      <div className="d-flex flex-row justify-content-center align-items-center">
        <IconButton className="mr-3" onClick={setPreviousMonth}>
          <KeyboardArrowLeftIcon fontSize="large" />
        </IconButton>
        <Typography>{currentMonthViewed.format("MMM YYYY")}</Typography>
        <IconButton className="ml-3" onClick={setNextMonth}>
          <KeyboardArrowRightIcon fontSize="large" />
        </IconButton>
      </div>
      <Chart
        chartType="PieChart"
        width="100%"
        height="400px"
        data={formattedCategoryGraphData}
        options={categoryGraphOptions}
      />
      <Chart
        chartType="PieChart"
        width="100%"
        height="400px"
        data={formattedAccountGraphData}
        options={accountGraphOptions}
      />
    </div>
  );
}
