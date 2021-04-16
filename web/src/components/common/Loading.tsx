
import { Overlay, Spinner, SpinnerSize, Stack } from "@fluentui/react";
import React from "react";
import { useTranslation } from "react-i18next";

export interface ILoadingProps {
  loading: boolean
}

export const Loading = (props: ILoadingProps) => {
  const { loading } = props, { t } = useTranslation();

  return (<>
    {loading &&
      <Overlay styles={{ root: { zIndex: 1 } }}>
        <Stack style={{ height: '100%' }} horizontalAlign="center" verticalAlign="center">
          <Spinner size={SpinnerSize.large} label={t('Load...')} />
        </Stack>
      </Overlay>
    }
  </>)
}