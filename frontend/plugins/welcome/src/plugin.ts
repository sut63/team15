import { createPlugin } from '@backstage/core';
import CreateDeposit from './components/RecordDeposit';
import ShowDeposit from './components/DepositTable';
import SignIn from './components/SignIn';
import CreateRoomdetail from './components/RoomDetails';
import WelcomePage from './components/WelcomePage';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', SignIn);
	  router.registerRoute('/DormEmployee', WelcomePage);
    router.registerRoute('/DepositTable', ShowDeposit);
    router.registerRoute('/RecordDeposit', CreateDeposit);
    router.registerRoute('/RoomDetails', CreateRoomdetail);

  },
});
