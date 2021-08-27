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
            <Text>Version: </Text>
            <Text>{aboutData?.version}</Text>
          </Stack>
          <Stack horizontal horizontalAlign="space-between">
            <Text>Commit: </Text>
            <Text>{aboutData?.commit}</Text>
          </Stack>
          <Stack horizontal horizontalAlign="space-between">
            <Text>Date: </Text>
            <Text>{aboutData?.date}</Text>
          </Stack>
          <Stack horizontal horizontalAlign="space-between">
            <Text>BuiltBy: </Text>
            <Text>{aboutData?.builtBy}</Text>
          </Stack>
          <Stack horizontal horizontalAlign="space-between">
            <Text>Environment: </Text>
            <Text>{aboutData?.environment}</Text>
          </Stack>
        </Stack>
      </Stack>
    </Dialog>
  </>)

}