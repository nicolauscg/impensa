import React, { useState, useEffect } from "react";
import * as R from "ramda";
import MUIDataTable from "mui-datatables";
import InfiniteScroll from "react-infinite-scroller";

import { objFromListWith } from "../../ramdaHelpers";
import useDebounce from "../../helper/customHooks";

const moment = require("moment");

const dataTableFormatHelper = {
  createIdToDataMap: objFromListWith(R.prop("id"))
};

export const DataTableFormatter = {
  formatDateFrom: keyFunc =>
    R.pipe(
      keyFunc,
      dateIsoString => moment(dateIsoString).format("DD/MM/YYYY")
    ),
  mapFromLookup: (lookupKeyFunc, lookupData, displayDataFunc) => dataRow =>
    R.applyTo(
      dataTableFormatHelper.createIdToDataMap(lookupData),
      R.pipe(
        R.prop(lookupKeyFunc(dataRow)),
        displayDataFunc
      )
    )
};

export default function MuiDataTable({
  headerNames,
  dataFormatters,
  data,
  onEdit,
  onDelete,
  onSearch,
  categoriesData,
  onLoadMore,
  hasMore
}) {
  const [muiTableData, setMuiTableData] = useState([]);
  const [searchQuery, setSearchQuery] = useState(undefined);
  const debouncedSearchQuery = useDebounce(searchQuery, 500);

  useEffect(() => {
    const formattedData = data.map(row => {
      const formattedRow = Object.assign({}, row);
      for (const [key, dataFormatter] of Object.entries(dataFormatters)) {
        formattedRow[key] = dataFormatter(row);
      }

      return formattedRow;
    });
    setMuiTableData(formattedData);
  }, [data, categoriesData]);

  useEffect(() => {
    if (debouncedSearchQuery !== undefined) {
      onSearch(debouncedSearchQuery);
    }
  }, [debouncedSearchQuery]);

  const optionsMui = {
    // still todo
    filter: false,
    onRowClick: (rowData, rowMeta) => onEdit(data[rowMeta.dataIndex]),
    filterType: "checkbox",
    print: false,
    download: false,
    onRowsDelete: ({ data: rowsDataToDelete }) => {
      const ids = rowsDataToDelete.map(el => muiTableData[el.index].id);

      return onDelete(ids);
    },
    serverSide: true,
    onSearchChange: queryString => {
      setSearchQuery(queryString);
    },
    onSearchClose: () => {
      onSearch("");
    },
    searchText: searchQuery,
    customFooter: () => ""
  };

  return (
    <InfiniteScroll loadMore={onLoadMore} hasMore={hasMore}>
      <MUIDataTable
        data={muiTableData}
        columns={headerNames}
        options={optionsMui}
      />
    </InfiniteScroll>
  );
}
