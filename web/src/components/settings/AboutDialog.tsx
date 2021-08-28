import { Dialog, Image, Stack, Text } from '@fluentui/react';
import React, { useEffect, useState } from 'react';
import logo from '../../logo.svg';
import { About, about } from '../../services/config.service';

export interface IAboutDialogProps {
  hidden: boolean
  onDismiss: () => void
}

export const AboutDialog = (props: IAboutDialogProps) => {

  const [aboutData, setAboutData] = useState<About>();

  useEffect(() => {
    about().then((ret) => {
      if (!ret) return;
      setAboutData(ret)
    })
      .finally(() => { })
  }, []);

  return (<>
    <Dialog
      minWidth={450}
      hidden={props.hidden}
      onDismiss={props.onDismiss}
      dialogContentProps={{
        title: "Redis Web Manager"
      }}
      modalProps={{ isBlocking: true }}>
      <Stack horizontal tokens={{ childrenGap: 20 }}>
        <Stack horizontalAlign="center" tokens={{ childrenGap: 5 }}>
          <Image src={logo} width={80} />
          <a href="https://github.com/slowrookie/redis-web-manager" target="_blank" rel="noreferrer">
            <Image src="https://img.shields.io/github/stars/slowrookie/redis-web-manager?logo=github" />
          </a>
        </Stack>

        <Stack grow={1} tokens={{ childrenGap: 5 }}>
          <Stack horizontal horizontalAlign="space-between">
            <Text variant="small">Version: </Text>
            <Text variant="small">{aboutData?.version}</Text>
          </Stack>
          <Stack horizontal horizontalAlign="space-between">
            <Text variant="small">Commit: </Text>
            <Text variant="small">{aboutData?.commit}</Text>
          </Stack>
          <Stack horizontal horizontalAlign="space-between">
            <Text variant="small">Date: </Text>
            <Text variant="small">{aboutData?.date}</Text>
          </Stack>
          <Stack horizontal horizontalAlign="space-between">
            <Text variant="small">BuiltBy: </Text>
            <Text variant="small">{aboutData?.builtBy}</Text>
          </Stack>
          <Stack horizontal horizontalAlign="space-between">
            <Text variant="small">Environment: </Text>
            <Text variant="small">{aboutData?.environment}</Text>
          </Stack>
        </Stack>
      </Stack>
    </Dialog>
  </>)

}