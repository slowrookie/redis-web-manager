import { DefaultButton, Dialog, DialogFooter, PrimaryButton, TextField } from '@fluentui/react';
import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { checkPort } from '../../services/config.service';

export interface IAboutDialogProps {
  port: number
  hidden: boolean
  onDismiss: () => void
  onChangePort: (port: number) => void
}

export const PortDialog = (props: IAboutDialogProps) => {

  const { t } = useTranslation(),
    [port, setPort] = useState<number>(props.port),
    [hasError, setHasError] = useState<boolean>(false);

  const handleDismiss = () => {
    props.onDismiss();
    setPort(props.port);
  }

  return (<>
    <Dialog
      hidden={props.hidden}
      onDismiss={handleDismiss}
      dialogContentProps={{
        title: t("Change port")
      }}
      modalProps={{ isBlocking: true }}>

      <TextField type="number" required underlined description={t("After the port is changed, it needs to be restarted to take effect")} value={`${port}`}
        onGetErrorMessage={(v) => {
          setHasError(false);
          if (Number(v) === props.port) {
            return
          }
          if (Number(v) <= 0 || Number(v) > 65536) {
            setHasError(true);
            return t('Port valid error');
          }
          return checkPort(Number(v))
            .then(() => "")
            .catch(err => {
              setHasError(true);
              return err.response.data || err.message;
            })
        }}
        onChange={(e, v) => { setPort(Number(v)) }}
      />

      <DialogFooter>
        <PrimaryButton disabled={hasError || !port || port === props.port} onClick={() => {
          props.onChangePort(port);
          handleDismiss();
        }} text={t('OK')} />
        <DefaultButton onClick={handleDismiss} text={t('Cancel')} />
      </DialogFooter>
    </Dialog>
  </>)

}