package edu.umflint.superteam.compromise;

import android.content.Intent;
import android.content.SharedPreferences;
import android.graphics.Color;
import android.os.AsyncTask;
import android.os.Bundle;
import android.preference.PreferenceManager;
import android.support.constraint.solver.ArrayLinkedVariables;
import android.support.design.widget.FloatingActionButton;
import android.support.design.widget.Snackbar;
import android.support.v4.content.ContextCompat;
import android.util.Log;
import android.view.SubMenu;
import android.view.View;
import android.support.design.widget.NavigationView;
import android.support.v4.view.GravityCompat;
import android.support.v4.widget.DrawerLayout;
import android.support.v7.app.ActionBarDrawerToggle;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.ArrayAdapter;
import android.widget.EditText;
import android.widget.ExpandableListAdapter;
import android.widget.ExpandableListView;
import android.widget.ListView;
import android.widget.TextView;
import android.widget.Toast;

import com.aurelhubert.ahbottomnavigation.AHBottomNavigation;
import com.aurelhubert.ahbottomnavigation.AHBottomNavigationAdapter;
import com.aurelhubert.ahbottomnavigation.AHBottomNavigationItem;
import com.aurelhubert.ahbottomnavigation.notification.AHNotification;

import org.json.JSONArray;
import org.json.JSONObject;

import java.io.BufferedReader;
import java.io.DataOutputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

import edu.umflint.superteam.compromise.API.GetRewards;
import edu.umflint.superteam.compromise.API.GetTasks;
import edu.umflint.superteam.compromise.Adapter.ExpandListAdapter;
import edu.umflint.superteam.compromise.Classes.ExpandListChild;
import edu.umflint.superteam.compromise.Classes.ExpandListGroup;


public class MainActivity extends AppCompatActivity {


    private ExpandListAdapter ExpAdapter;
    private ArrayList<ExpandListGroup> ExpListItems;
    private ExpandableListView ExpandList;

    public class UserFindGroups extends AsyncTask<Void, Void, JSONArray> {

        private final String mEmail;
        private StringBuffer response;

        UserFindGroups(String email) {
            mEmail = email;
        }

        @Override
        protected JSONArray doInBackground(Void... params) {
            // TODO: attempt authentication against a network service.

            try {
                String url = "http://api.compromise.rocks/api/groups/" + mEmail;
                URL obj = new URL(url);
                HttpURLConnection con = (HttpURLConnection) obj.openConnection();

                int responseCode = con.getResponseCode();
                Log.i("Groups", "Sending 'GET' request to URL : " + url);
                Log.i("Groups", "Response Code : " + responseCode);

                BufferedReader in = new BufferedReader(
                        new InputStreamReader(con.getInputStream()));
                String inputLine;
                response = new StringBuffer();

                while ((inputLine = in.readLine()) != null) {
                    response.append(inputLine);
                }
                in.close();
            } catch (Exception ex) {
                Log.e("Groups", ex.toString());
                return null;
            }

            //print result
            Log.i("Groups", response.toString());
            if (!response.toString().isEmpty()) {
                try {
                    JSONArray jsonObj = new JSONArray(response.toString());
                    return jsonObj;
                } catch (Exception ex) {
                    return null;
                }

            } else {
                return null;
            }
        }

        protected void onPostExecute(final JSONArray array) {
            int groupId = 0;
            String groupName = "";
            int points = 0;
            NavigationView nav = (NavigationView) findViewById(R.id.nav_view);
            Menu menu = nav.getMenu();

            try {
                if (array != null) {

                    for (int i = 0; i < array.length(); i++) {
                        JSONObject row = array.getJSONObject(i);
                        groupId = row.getInt("GroupId");
                        groupName = row.getString("GroupName");
                        points = row.getInt("TotalPoints");
                        menu.add(0, groupId, 0, "(" + points + " pts) " + groupName);
                    }

                } else {

                }
            } catch (Exception ex) {

            }
        }

    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        final Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
        setSupportActionBar(toolbar);


        AHBottomNavigation bottomNavigation = (AHBottomNavigation) findViewById(R.id.bottom_navigation);

        // Create items
        AHBottomNavigationItem item1 = new AHBottomNavigationItem(R.string.tasks, R.drawable.ic_tasks, R.color.colorPrimary);
        AHBottomNavigationItem item2 = new AHBottomNavigationItem(R.string.rewards, R.drawable.ic_rewards, R.color.colorPrimary);

        // Add items
        bottomNavigation.addItem(item1);
        bottomNavigation.addItem(item2);

        // Set background color
        bottomNavigation.setDefaultBackgroundColor(Color.parseColor("#FEFEFE"));

        // Disable the translation inside the CoordinatorLayout
        bottomNavigation.setBehaviorTranslationEnabled(false);

        // Change colors
        bottomNavigation.setAccentColor(Color.parseColor("#F63D2B"));
        bottomNavigation.setInactiveColor(Color.parseColor("#303F9F"));

        // Force to tint the drawable (useful for font with icon for example)
        bottomNavigation.setForceTint(true);

        // Force the titles to be displayed (against Material Design guidelines!)
        bottomNavigation.setForceTitlesDisplay(true);
        // Or force the titles to be hidden (against Material Design guidelines, too!)
        bottomNavigation.setForceTitlesHide(true);

        // Use colored navigation with circle reveal effect
        bottomNavigation.setColored(true);

        // Set current item programmatically
        bottomNavigation.setCurrentItem(0);

        // Set listeners
        bottomNavigation.setOnTabSelectedListener(new AHBottomNavigation.OnTabSelectedListener() {
            @Override
            public boolean onTabSelected(int position, boolean wasSelected) {
                if(position == 0)
                {

                } else if (position == 1) {

                }
                return true;
            }
        });


        SharedPreferences prefs = PreferenceManager.getDefaultSharedPreferences(getApplicationContext());
        UserFindGroups groupTask = new UserFindGroups(prefs.getString("EmailAddress", ""));
        groupTask.execute((Void) null);
        NavigationView navigationView = (NavigationView) findViewById(R.id.nav_view);

        NavigationView.OnNavigationItemSelectedListener item_click_listener = new NavigationView.OnNavigationItemSelectedListener() {
            @Override
            public boolean onNavigationItemSelected(MenuItem item) {

                int id = item.getItemId();
                SharedPreferences prefs = PreferenceManager.getDefaultSharedPreferences(getApplicationContext());
                SharedPreferences.Editor editor = prefs.edit();
                editor.putInt("SelectedGroup", id);
                editor.apply();

                toolbar.setTitle(item.getTitle());

                ExpandList = (ExpandableListView) findViewById(R.id.list);

                ExpListItems = SetStandardGroups(id);
                ExpAdapter = new ExpandListAdapter(MainActivity.this, ExpListItems);
                ExpandList.setAdapter(ExpAdapter);

                ExpandList.setOnChildClickListener(new ExpandableListView.OnChildClickListener() {

                    @Override
                    public boolean onChildClick(ExpandableListView parent, View v, int groupPosition, int childPosition, long id) {
                        Log.i("Clicked", "Item Clicked!");
                        // TODO Auto-generated method stub
                        ExpandListChild selected = (ExpandListChild) ExpAdapter.getChild(groupPosition, childPosition);
                        //Toast.makeText(getApplication().getBaseContext(), "Item Clicked! " + selected.getTag(), Toast.LENGTH_SHORT).show();
                        Intent intent = new Intent(getApplicationContext(), DetailedTaskActivity.class);
                        intent.putExtra("taskTitle", selected.getName());
                        intent.putExtra("taskPoints", selected.getPoints());
                        intent.putExtra("taskDescription", selected.getDescription());
                        intent.putExtra("newTask", false);
                        startActivity(intent);
                        Log.i("Object", selected.toString());
                        return true;
                    }
                });

                DrawerLayout drawer = (DrawerLayout) findViewById(R.id.drawer_layout);
                drawer.closeDrawer(GravityCompat.START);
                return true;
            }
        };
        navigationView.setNavigationItemSelectedListener(item_click_listener);

        final DrawerLayout drawer = (DrawerLayout) findViewById(R.id.drawer_layout);
        ActionBarDrawerToggle toggle = new ActionBarDrawerToggle(
                this, drawer, toolbar, R.string.navigation_drawer_open, R.string.navigation_drawer_close);
        drawer.setDrawerListener(toggle);
        toggle.syncState();
    }

    private ArrayList<ExpandListGroup> SetStandardGroups(int groupId)
    {
        AHBottomNavigation bottomNavigation = (AHBottomNavigation) findViewById(R.id.bottom_navigation);
        ArrayList list = new ArrayList();
        try {
            if(bottomNavigation.getCurrentItem() == 0)
                list = new GetTasks(groupId).execute().get();
            else
                list = new GetRewards(groupId).execute().get();
        } catch (Exception e) {

        }
        return list;
    }

    @Override
    public void onBackPressed() {
        DrawerLayout drawer = (DrawerLayout) findViewById(R.id.drawer_layout);
        if (drawer.isDrawerOpen(GravityCompat.START)) {
            drawer.closeDrawer(GravityCompat.START);
        } else {
            super.onBackPressed();
        }
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.main, menu);

        TextView loggedInAccount = (TextView) findViewById(R.id.loggedInAccount);
        SharedPreferences prefs = PreferenceManager.getDefaultSharedPreferences(getApplicationContext());
        loggedInAccount.setText(prefs.getString("EmailAddress", ""));


        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();

        //noinspection SimplifiableIfStatement
        if (id == R.id.logout) {
            SharedPreferences prefs = PreferenceManager.getDefaultSharedPreferences(getApplicationContext());
            SharedPreferences.Editor editor = prefs.edit();
            editor.putString("EmailAddress", "");
            editor.putInt("SelectedGroup", -1);
            editor.apply();
            Toast.makeText(getApplication().getBaseContext(), "You have been logged out.", Toast.LENGTH_SHORT).show();
            Intent intent = new Intent(this, LoginActivity.class);
            startActivity(intent);

        } else if(id == R.id.action_group_directory)
        {
            Intent intent = new Intent(this, GroupMemberActivity.class);
            startActivity(intent);
        } else if(id == R.id.action_new)
        {
            Intent intent = new Intent(this, DetailedTaskActivity.class);
            intent.putExtra("taskTitle", "Example Title");
            intent.putExtra("taskPoints", 999);
            intent.putExtra("taskDescription", "Example Description");
            intent.putExtra("newTask", true);
            startActivity(intent);
        }

        return super.onOptionsItemSelected(item);
    }
}
