package com.example.unittest;

import android.util.Log;
import android.widget.EditText;
import android.widget.Toast;

import org.junit.After;
import org.junit.Before;
import org.junit.Rule;
import org.junit.Test;

import androidx.test.espresso.Espresso;
import androidx.test.rule.ActivityTestRule;

import static androidx.test.espresso.action.ViewActions.click;
import static androidx.test.espresso.matcher.ViewMatchers.withId;


public class MainActivityTest {

    private EditText ed_textone;

    @Rule
    public ActivityTestRule<MainActivity> mainActivityActivityTestRule = new ActivityTestRule<>(MainActivity.class);
    private MainActivity mainActivity = null;

    @Before
    public void setUp() throws Exception {
        try {
            mainActivity = mainActivityActivityTestRule.getActivity();
            ed_textone = mainActivity.findViewById(R.id.ed_Textone);
            System.out.println("all is not null");
        }
    catch (NullPointerException e)
        {
            if (mainActivity == null) System.out.println("mainActivity is null");
        }

    }

    @After
    public void tearDown() throws Exception {
    }

    @Test
    public void onCreate() {
        Log.e("tag:oncreate","oncreate()");
    }

    @Test
    public void onJump() {
        Log.e("tag:onjump","jump()");
        ed_textone.setText("123");
        if(ed_textone.getText().toString() == null)
        {
            System.out.println("edit is nul");
        }
        Toast.makeText(mainActivity.getApplication(),ed_textone.getText().toString(), Toast.LENGTH_SHORT).show();
        mainActivity.findViewById(R.id.btn_jump);// perform the button click
        Espresso.onView(withId(R.id.btn_jump)).perform(click());
        }

}