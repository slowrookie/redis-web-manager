import { Dropdown, IconButton, IDropdownOption, Stack, Text, TooltipHost } from "@fluentui/react";
import React, { FormEvent, useEffect, useState } from "react";
import { useTranslation } from "react-i18next";

const PageSizeOptions: Array<{ key: number, text: string }> = [
  { key: 10, text: '10' },
  { key: 20, text: '20' },
  { key: 50, text: '50' },
  { key: 100, text: '100' },
  { key: 500, text: '500' },
  { key: 1000, text: '1000' },
  { key: 2000, text: '2000' },
]

export interface IPaginationProps {
  onChange: (start: number, end: number) => void
  totalElements: number
  defaultPageSize?: number
}

export interface IPagination {
  empty: boolean
  first: boolean
  last: boolean
  pageNumber: number
  pageSize: number
  loadedElements: number
  totalElements: number
  totalPages: number
  start: number
  end: number
}

const defaultPagination: IPagination = {
  empty: true,
  first: false,
  last: false,
  pageNumber: 1,
  pageSize: 20,
  loadedElements: 0,
  totalElements: 0,
  totalPages: 0,
  start: 0,
  end: 19
}

export const Pagination = (props: IPaginationProps) => {

  const { onChange, totalElements, defaultPageSize } = props,
    { t } = useTranslation();

  const [pagination, setPagination] = useState({ ...defaultPagination, pageSize: defaultPageSize || 20 }),
    [pages, setPages] = useState<Array<IDropdownOption>>([]);

  useEffect(() => {
    const pageSize = pagination.pageSize;
    const totalPages = Math.ceil((totalElements || 0) / (pageSize));
    setPagination(page => {
      return {
        ...page,
        empty: !totalElements,
        first: pagination.pageNumber === 1,
        last: pagination.pageNumber === totalPages,
        pageNumber: pagination.pageNumber,
        pageSize,
        totalElements: totalElements,
        totalPages: totalPages,
        start: (pagination.pageNumber - 1) * pageSize,
        end: (pagination.pageNumber - 1) * pageSize + pageSize - 1
      }
    });
  }, [totalElements, pagination.pageSize, pagination.pageNumber]);

  useEffect(() => {
    setPages([...Array(pagination.totalPages)].map((n, i) => { return { key: i + 1 + '', text: i + 1 + '' } }));
  }, [pagination.totalPages])

  useEffect(() => {
    onChange && onChange(pagination.start, pagination.end);
    // eslint-disable-next-line
  }, [pagination.start, pagination.end])

  const handlePageSizeChange = (e: FormEvent, o?: IDropdownOption) => {
    o && setPagination({ ...pagination, pageSize: Number(o.key), pageNumber: 1 });
  }

  const handlePrePage = () => {
    setPagination({ ...pagination, pageNumber: pagination.pageNumber - 1 });
  }

  const handleNextPage = () => {
    setPagination({ ...pagination, pageNumber: pagination.pageNumber + 1 });
  }

  const handlePageNumberChange = (e: FormEvent, o?: IDropdownOption) => {
    o && setPagination({ ...pagination, pageNumber: Number(o.key) });
  }

  return (
    <Stack horizontal horizontalAlign="end" verticalAlign="center" tokens={{ childrenGap: 10, padding: 5 }}>

      <Dropdown
        selectedKey={pagination.pageSize}
        options={PageSizeOptions}
        onChange={handlePageSizeChange}
      />
      <Text variant="smallPlus" >{`${t('items')}/${t('page')}`}</Text>

      <Text variant="smallPlus" >{`${t('total')} ${pagination.totalElements} ${t('items')}`}</Text>

      <TooltipHost content={t('Previous page')}>
        <IconButton iconProps={{ iconName: 'ChevronLeft' }} disabled={pagination.empty || pagination.first}
          onClick={handlePrePage} />
      </TooltipHost>

      <TooltipHost content={t('Next page')}>
        <IconButton iconProps={{ iconName: 'ChevronRight' }} disabled={pagination.empty || pagination.last}
          onClick={handleNextPage} />
      </TooltipHost>

      <Dropdown
        selectedKey={pagination.pageNumber}
        options={pages}
        onChange={handlePageNumberChange}
      />
      <Text variant="smallPlus" >{`/ ${pagination.totalPages} ${t('page')}`}</Text>

    </Stack>
  );
}