import React from "react";
import { Chart } from "react-google-charts";
import {
  useAxiosSafely,
  urlGraphTransactionCategory,
  urlGraphTransactionAccount
} from "../../api";

export default function GraphPage() {
  const [{ data: categoryGraphData }] = useAxiosSafely(
    urlGraphTransactionCategory()
  );
  const [{ data: accountGraphData }] = useAxiosSafely(
    urlGraphTransactionAccount()
  );

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
